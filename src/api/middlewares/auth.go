package middlewares

import (
	"context"
	"errors"
	"github.com/alireza-fa/blog-go/src/api/helper"
	"github.com/alireza-fa/blog-go/src/constants"
	"github.com/alireza-fa/blog-go/src/services"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func Authentication(next http.Handler) http.Handler {
	var tokenService = services.NewTokenService()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		claimMap := map[string]interface{}{}
		auth := r.Header.Get(constants.AuthorizationHeaderKey)
		token := strings.Split(auth, " ")
		if auth == "" {
			err = errors.New("token is required")
		} else {
			claimMap, err = tokenService.GetClaims(token[1])
			if err != nil {
				switch {
				case errors.Is(err, jwt.ErrTokenExpired):
					err = errors.New("token expired")
				default:
					err = errors.New("token is invalid")
				}
			}
		}
		if err != nil {
			helper.BaseResponseWithError(w, nil, http.StatusUnauthorized, err)
			return
		}

		ctx := context.WithValue(r.Context(), constants.UserIdKey, claimMap[constants.UserIdKey])
		ctx = context.WithValue(ctx, constants.FullNameKey, claimMap[constants.UserNameKey])
		ctx = context.WithValue(ctx, constants.UserNameKey, claimMap[constants.UserNameKey])
		ctx = context.WithValue(ctx, constants.EmailKey, claimMap[constants.EmailKey])
		ctx = context.WithValue(ctx, constants.RolesKey, claimMap[constants.RolesKey])
		ctx = context.WithValue(ctx, constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
