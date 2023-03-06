package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	atData	= "accessToken"
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