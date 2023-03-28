package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"

	model "github.com/vnSasa/music-market-api/model"
)

func (h *Handler) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (h *Handler) initAdmin() {
	input := model.User{
		Login:    viper.GetString("admin.Login"),
		Password: viper.GetString("admin.Password"),
	}
	h.services.Authorization.CreateUser(input)
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
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     atData,
		Value:    token.AccessToken,
		Path:     "/",
		Expires:  time.Unix(token.AtExpires, 0),
		HttpOnly: true,
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     rtData,
		Value:    token.RefreshToken,
		Path:     "/",
		Expires:  time.Unix(token.RtExpires, 0),
		HttpOnly: true,
	})

	if strings.Compare(input.Login, viper.GetString("admin.Login")) == 0 {
		c.HTML(http.StatusOK, "main_page_admin.html", nil)
	} else {
		c.HTML(http.StatusOK, "main_page_user.html", nil)
	}
}

func (h *Handler) logout(c *gin.Context) {
	accessTokenValue, err := h.getAccessToken(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "access token not found")

		return
	}
	_, err = h.services.Authorization.ParseToken(accessTokenValue)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     atData,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     rtData,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})
	c.HTML(http.StatusOK, "index.html", nil)
}
