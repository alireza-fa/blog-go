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

func (service *OtpService) SetOtp(user dto.CreateUser) error {
	key := user.UserName + "_otp"

	code, _ := cache.Get[int](service.redisClient, key)

	if code != 0 {
		service.logger.Warn(logging.Otp, logging.OtpGenerate,
			"you can get just one code in two minutes", map[logging.ExtraKey]interface{}{logging.UserName: user.UserName, logging.Email: user.Email})
		return errors.New("you can get just one code in two minutes")
	}

	code = rand.Intn(9999-999) + 999

	err := cache.Set[int](service.redisClient, key, code, time.Second*120)
	if err != nil {
		return err
	}

	extraInfo := map[logging.ExtraKey]interface{}{logging.UserName: user.UserName, logging.Email: user.Email}
	service.logger.Info(logging.Redis, logging.RedisSet, fmt.Sprintf("Set a otp code for %s code: %d", user.UserName, code), extraInfo)

	return nil
}

func (service *OtpService) GetOtp(key string) {}
