package services

import (
	"github.com/alireza-fa/blog-go/src/api/dto"
	"github.com/alireza-fa/blog-go/src/data/cache"
	"github.com/alireza-fa/blog-go/src/pkg/logging"
	"github.com/go-redis/redis"
)

type UserFrontService struct {
	logger      logging.Logger
	otpService  *OtpService
	redisClient *redis.Client
}

func NewUserFrontService() *UserFrontService {
	return &UserFrontService{
		logger:      logging.NewLogger(),
		otpService:  NewOtpService(),
		redisClient: cache.GetRedis(),
	}
}

func (service *UserFrontService) CreateUser(user dto.CreateUser) error {
	return service.otpService.SetOtp(user)
}
