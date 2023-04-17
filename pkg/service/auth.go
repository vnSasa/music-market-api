package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vnSasa/music-market-api/model"
	"github.com/vnSasa/music-market-api/pkg/repository"
)

const (
	signingKey          = "qrkjk#4#%35FSFJlja#4353KSFjH"
	salt                = "hjqrhjqw124617ajfhajs"
	timeForAccessToken  = 1
	timeForRefreshToken = 24 * 7
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) error {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUserByID(id int) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *AuthService) UpdateUser(userID int, input model.User) error {
	if input.Password != "" {
		input.Password = generatePasswordHash(input.Password)

		return s.repo.UpdateUserWithPassword(userID, input)
	} else {
		return s.repo.UpdateUserWithoutPassword(userID, input)
	}

}

func (s *AuthService) GenerateToken(login, password string) (*model.TokenDetails, string, error) {
	id, err := s.repo.GetUserID(login, generatePasswordHash(password))
	if err != nil {
		return nil, "", err
	}

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, "", err
	}

	td := &model.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * timeForAccessToken).Unix()

	td.RtExpires = time.Now().Add(time.Hour * timeForRefreshToken).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.AtExpires,
		},
		UserID:  id,
		Status: user.Status,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.RtExpires,
		},
		UserID:    id,
		Status: user.Status,
		IsRefresh: true,
	})

	td.AccessToken, err = accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, "", err
	}

	td.RefreshToken, err = refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, "", err
	}

	return td, user.Status, nil
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

func (s *AuthService) RefreshToken(refreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &model.RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*model.RefreshTokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	if !claims.IsRefresh {
		return "", errors.New("is not refresh token")
	}

	atExpires := time.Now().Add(time.Minute * timeForAccessToken).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: atExpires,
		},
		UserID:  claims.UserID,
		Status: claims.Status,
	})

	accessTokenValue, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return accessTokenValue, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
