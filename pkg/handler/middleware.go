package handler

import (
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	atData = "accessToken"
	rtData = "refreshToken"
)

func saveAccessToken(c *gin.Context) {
	accessTokenValue, err := c.Cookie(atData)
	if err != nil {
		c.Set(atData, "")
	} else {
		c.Set(atData, accessTokenValue)
	}
	c.Next()
}

func (h *Handler) getAccessToken(c *gin.Context) (string, error) {
	accessTokenValue, err := c.Cookie(atData)
	if err != nil {
		refreshTokenValue, err := c.Cookie(rtData)
		if err != nil {
			return "", errors.New("token not found")
		}

		accessTokenValue, err = h.services.Authorization.RefreshToken(refreshTokenValue)
		if err != nil {
			return "", errors.New(err.Error())
		}
	}

	return accessTokenValue, nil
}

func (h *Handler) adminIdentity(c *gin.Context) {
	accessTokenValue, err := h.getAccessToken(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	accessToken, err := h.services.Authorization.ParseToken(accessTokenValue)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	if !accessToken.IsAdmin {
		newErrorResponse(c, http.StatusInternalServerError, "only admin have access")

		return
	}
}
