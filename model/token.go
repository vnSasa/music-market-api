package model

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

type AccessTokenClaims struct {
	jwt.StandardClaims
	UserID  int  `json:"user_id"`
	Status string `json:"status"`
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
	UserID    int  `json:"user_id"`
	Status   string `json:"status"`
	IsRefresh bool `json:"is_refresh"`
}
