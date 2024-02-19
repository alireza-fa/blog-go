package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/alireza-fa/blog-go/src/api/dto"
	"github.com/alireza-fa/blog-go/src/constants"
	"github.com/alireza-fa/blog-go/src/data/cache"
	"github.com/alireza-fa/blog-go/src/data/db"
	"github.com/alireza-fa/blog-go/src/data/models"
	"github.com/alireza-fa/blog-go/src/pkg/logging"
	"github.com/alireza-fa/blog-go/src/pkg/notification"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserFrontService struct {
	logger       logging.Logger
	otpService   *OtpService
	redisClient  *redis.Client
	notification notification.Notifier
	tokenService *TokenService
	database     *gorm.DB
}

func NewUserFrontService() *UserFrontService {
	return &UserFrontService{
		logger:       logging.NewLogger(),
		otpService:   NewOtpService(),
		redisClient:  cache.GetRedis(),
		notification: notification.NewNotifier(),
		tokenService: NewTokenService(),
		database:     db.GetDb(),
	}
}

func (service *UserFrontService) CreateUser(user dto.CreateUser) error {
	exist, err := service.checkUserNameExist(user.UserName)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("this username already exists")
	}

	exist, err = service.checkEmailExist(user.Email)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("this email already taken")
	}

	key := user.UserName + "info"

	code, err := service.otpService.SetOtp(user)
	if err != nil {
		return err
	}

	err = cache.Set[dto.CreateUser](service.redisClient, key, user, time.Second*120)
	if err != nil {
		return err
	}
	service.logger.Info(logging.Redis, logging.RedisSet, "user information set on cache for 120 seconds",
		map[logging.ExtraKey]interface{}{logging.UserName: user.UserName, logging.Email: user.Email})

	service.notification.Send(user.Email, map[string]string{notification.Message: strconv.Itoa(code)})

	return nil
}

func (service *UserFrontService) VerifyUser(userVerify *dto.UserVerify) (*models.User, error) {
	var userCreate dto.CreateUser

	_, err := service.otpService.VerifyOtp(userVerify)
	if err != nil {
		return nil, err
	}

	key := userVerify.UserName + "info"
	userCreate, err = cache.Get[dto.CreateUser](service.redisClient, key)
	if err != nil {
		service.logger.Error(logging.Redis, logging.RedisGet, fmt.Sprintf("can't create user %s", userVerify.UserName), nil)
		return nil, err
	}

	var user models.User = models.User{
		UserName: userCreate.UserName,
		Email:    userCreate.Email,
		FullName: userCreate.FullName,
	}

	bytePassword, err := bcrypt.GenerateFromPassword([]byte(userCreate.Password), bcrypt.DefaultCost)
	if err != nil {
		service.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
	}

	user.Password = string(bytePassword)

	database := service.database.Begin()

	if err := database.Create(&user).Error; err != nil {
		database.Rollback()
		service.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
	}
	var defaultRole models.Role
	if err := database.
		Model(models.Role{}).
		Where("name = ?", constants.DefaultRole).
		Find(&defaultRole).Error; err != nil {
		database.Rollback()
		service.logger.Error(logging.Postgres, logging.Rollback, "default role not found"+err.Error(), nil)
		return nil, err
	}
	userRole := models.UserRole{UserId: user.Id, RoleId: defaultRole.Id}
	database.Create(&userRole)

	database.Commit()

	return &user, nil
}

func (service *UserFrontService) UserLogin(userLogin dto.UserLogin) (*dto.TokenDetail, error) {
	var user models.User

	if err := service.database.
		Model(&models.User{}).
		Where("user_name = ?", userLogin.UserName).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error; err != nil {
		return nil, err
	}

	if user.UserName == "" {
		return nil, errors.New("user with this information not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return nil, err
	}

	tokenDto := TokenDto{
		UserId:   user.Id,
		FullName: user.FullName,
		UserName: user.UserName,
		Email:    user.Email,
	}

	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tokenDto.Roles = append(tokenDto.Roles, ur.Role.Name)
		}
	}

	return service.tokenService.GenerateToken(&tokenDto)
}

func (service *UserFrontService) checkUserNameExist(username string) (bool, error) {
	var exists bool
	if err := service.database.Model(models.User{}).
		Select("count(*) > 0").
		Where("user_name = ?", username).
		Find(&exists).Error; err != nil {
		service.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (service *UserFrontService) checkEmailExist(email string) (bool, error) {
	var exists bool
	if err := service.database.Model(models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).Error; err != nil {
		service.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (service *UserFrontService) UserProfile(ctx context.Context) *dto.Profile {
	var userProfile dto.Profile

	userProfile.UserName = ctx.Value(constants.UserNameKey).(string)
	userProfile.FullName = ctx.Value(constants.FullNameKey).(string)
	userProfile.Email = ctx.Value(constants.EmailKey).(string)

	return &userProfile
}

func (service *UserFrontService) UserProfileUpdate(ctx context.Context, profileUpdate dto.ProfileUpdate) (*dto.Profile, error) {
	var user models.User
	user.FullName = profileUpdate.FullName

	if err := service.database.
		Model(models.User{}).
		Where("id = ?", ctx.Value(constants.UserIdKey)).
		Updates(&user).Error; err != nil {
		service.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}

	return service.UserProfile(ctx), nil
}
