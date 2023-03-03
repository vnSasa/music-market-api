package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"errors"
	model "github.com/vnSasa/music-market-api/model"
)

func (h *Handler) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (h *Handler) signUp(c *gin.Context) {
	c.HTML(http.StatusOK, "sign-up.html", nil)
}

func (h *Handler) newSignUp(c *gin.Context) {
	var input model.User
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "sign-up.html", "invalid input body")

		return
	}
	err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.HTML(http.StatusOK, "index.html", nil)
}

func (h *Handler) signIn(c *gin.Context) {
	var input model.SignInData
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "index.html", errors.New("invalid input body"))

		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		c.HTML(http.StatusBadRequest, "sign-up.html", errors.New("user not valid"))

		return
	}
	red := model.GetRedisConn()
	at := time.Unix(token.AtExpires, 0)
	rt := time.Unix(token.RtExpires, 0)
	now := time.Now()
	_, err = red.Set(c, token.AccessUUID, token.AccessToken, at.Sub(now)).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	_, err = red.Set(c, token.RefreshUUID, token.RefreshToken, rt.Sub(now)).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.HTML(http.StatusOK, "main_page.html", nil)
}