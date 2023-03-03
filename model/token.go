package model

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

type AccessTokenClaims struct {
	jwt.StandardClaims
	UserID  int    `json:"user_id"`
	AtUUID  string `json:"access_uuid"`
	RtUUID  string `json:"refresh_uuid"`
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
	UserID    int    `json:"user_id"`
	RtUUID    string `json:"refresh_uuid"`
	AtUUID    string `json:"access_uuid"`
	IsRefresh bool   `json:"is_refresh"`
}