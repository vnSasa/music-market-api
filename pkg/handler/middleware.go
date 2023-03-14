package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	model "github.com/vnSasa/music-market-api/model"
)

const (
	atData = "accessToken"
)

func (h *Handler) saveAccessToken(c *gin.Context) {
	accessTokenValue, err := c.Cookie(atData)
	if err != nil {
		c.Set(atData, "")
	} else {
		c.Set(atData, accessTokenValue)
	}
	c.Next()
}

func (h *Handler) adminIdentity(c *gin.Context) {
	accessTokenValue, err := c.Cookie(atData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "access token not found")

		return
	}
	accessToken, err := h.services.Authorization.ParseToken(accessTokenValue)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	red := model.GetRedisConn()
	_, err = red.Get(c, accessToken.AtUUID).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	if !accessToken.IsAdmin {
		newErrorResponse(c, http.StatusInternalServerError, "only admin have access")

		return
	}

}
