package services

import (
	"github.com/alireza-fa/blog-go/src/api/dto"
	"github.com/alireza-fa/blog-go/src/constants"
	"github.com/alireza-fa/blog-go/src/pkg/logging"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"
)

type TokenService struct {
	logger logging.Logger
}

type TokenDto struct {
	UserId   int
	FullName string
	UserName string
	Email    string
	Roles    []string
}

func NewTokenService() *TokenService {
	return &TokenService{
		logger: logging.NewLogger(),
	}
}

func (s *TokenService) GenerateToken(token *TokenDto) (*dto.TokenDetail, error) {
	td := &dto.TokenDetail{}

	accessTokenDuration, _ := strconv.Atoi(os.Getenv(constants.AccessTokenLifetime))
	td.AccessTokenExpireTime = time.Now().Add(time.Duration(accessTokenDuration) * time.Minute).Unix()

	refreshTokenDuration, _ := strconv.Atoi(os.Getenv(constants.RefreshTokenLifetime))
	td.RefreshTokenExpireTime = time.Now().Add(time.Duration(refreshTokenDuration) * time.Hour * 24).Unix()

	act := jwt.MapClaims{}

	act[constants.UserIdKey] = token.UserId
	act[constants.FullNameKey] = token.FullName
	act[constants.UserNameKey] = token.UserName
	act[constants.EmailKey] = token.Email
	act[constants.RolesKey] = token.Roles
	act[constants.ExpireTimeKey] = td.AccessTokenExpireTime
	ac := jwt.NewWithClaims(jwt.SigningMethodHS256, act)

	var err error
	td.AccessToken, err = ac.SignedString([]byte(os.Getenv(constants.AccessTokenSecretKey)))
	if err != nil {
		return nil, err
	}

	rtc := jwt.MapClaims{}

	rtc[constants.UserIdKey] = token.UserId
	rtc[constants.ExpireTimeKey] = td.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)

	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv(constants.RefreshTokenSecretKey)))
	if err != nil {
		return nil, err
	}

	return td, err
}
