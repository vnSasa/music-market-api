package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	
}