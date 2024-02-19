package dto

type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int64  `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int64  `json:"refreshTokenExpireTime"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type AccessToken struct {
	AccessToken string `json:"accessToken" validate:"required"`
}
