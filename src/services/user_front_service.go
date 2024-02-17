package services

import (
	"fmt"
	"github.com/alireza-fa/blog-go/src/api/dto"
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
	database     *gorm.DB
}

func NewUserFrontService() *UserFrontService {
	return &UserFrontService{
		logger:       logging.NewLogger(),
		otpService:   NewOtpService(),
		redisClient:  cache.GetRedis(),
		notification: notification.NewNotifier(),
		database:     db.GetDb(),
	}
}

func (service *UserFrontService) CreateUser(user dto.CreateUser) error {
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
	database.Commit()

	return &user, nil
}

func (service *UserFrontService) UserLogin(userLogin dto.UserLogin) (*models.User, error) {
	var user models.User

	if err := service.database.
		Model(&models.User{}).
		Where("user_name = ?", userLogin.UserName).
		Find(&user).Error; err != nil {
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}