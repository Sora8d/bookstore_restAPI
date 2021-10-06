package controller

import (
	"bookstoreapi/oauth/domain/access_token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByIdC(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func (handler *accessTokenHandler) GetByIdC(c *gin.Context) {
	accesTokenId := c.Param("access_token_id")
	accessToken, err := handler.service.GetById(accesTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}
