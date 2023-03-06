package service

import (
	"fmt"
	"crypto/sha256"
	"time"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"github.com/vnSasa/music-market-api/model"
	"github.com/vnSasa/music-market-api/pkg/repository"
)

const (
	tokenTTLup          = 10 * time.Minute
	signingKey          = "qrkjk#4#%35FSFJlja#4353KSFjH"
	salt = "hjqrhjqw124617ajfhajs"
	timeForAccessToken  = 15
	timeForRefreshToken = 24 * 7
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthSerice(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) error {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(login, password string) (*model.TokenDetails, error) {
	id, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return nil, err
	}
	td := &model.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * timeForAccessToken).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * timeForRefreshToken).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.AtExpires,
		},
		UserID:  id,
		AtUUID:  td.AccessUUID,
		RtUUID:  td.RefreshUUID,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.RtExpires,
		},
		UserID:    id,
		RtUUID:    td.RefreshUUID,
		AtUUID:    td.AccessUUID,
		IsRefresh: true,
	})

	td.AccessToken, err = accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}

	td.RefreshToken, err = refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *AuthService) ParseToken(accessToken string) (*model.AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &model.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.AccessTokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}