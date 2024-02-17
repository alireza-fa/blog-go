package services

import (
	"errors"
	"fmt"
	"github.com/alireza-fa/blog-go/src/api/dto"
	"github.com/alireza-fa/blog-go/src/data/cache"
	"github.com/alireza-fa/blog-go/src/pkg/logging"
	"github.com/go-redis/redis"
	"math/rand"
	"time"
)

type OtpService struct {
	logger      logging.Logger
	redisClient *redis.Client
}

func NewOtpService() *OtpService {
	return &OtpService{
		logger:      logging.NewLogger(),
		redisClient: cache.GetRedis(),
	}
}

func (service *OtpService) SetOtp(user dto.CreateUser) (int, error) {
	key := user.UserName + "_otp"

	code, _ := cache.Get[int](service.redisClient, key)

	if code != 0 {
		service.logger.Warn(logging.Otp, logging.OtpGenerate,
			"you can get just one code in two minutes", map[logging.ExtraKey]interface{}{logging.UserName: user.UserName, logging.Email: user.Email})
		return 0, errors.New("you can get just one code in two minutes")
	}

	code = rand.Intn(9999-999) + 999

	err := cache.Set[int](service.redisClient, key, code, time.Second*120)
	if err != nil {
		return 0, err
	}

	extraInfo := map[logging.ExtraKey]interface{}{logging.UserName: user.UserName, logging.Email: user.Email}
	service.logger.Info(logging.Redis, logging.RedisSet, fmt.Sprintf("Set a otp code for %s code: %d", user.UserName, code), extraInfo)

	return code, nil
}

func (service *OtpService) VerifyOtp(user *dto.UserVerify) (int, error) {
	key := user.UserName + "_otp"

	code, err := cache.Get[int](service.redisClient, key)
	if err != nil {
		service.logger.Info(logging.Otp, logging.OtpGet, fmt.Sprintf("not found otp code for %s", user.UserName),
			map[logging.ExtraKey]interface{}{logging.UserName: user.UserName, logging.OtpCode: user.Code})
		return code, err
	}

	if user.Code != code {
		service.logger.Info(logging.Otp, logging.OtpGet, fmt.Sprintf("not valid otp code %s", user.UserName), nil)
		return code, errors.New("invalid code")
	}

	return code, nil
}
