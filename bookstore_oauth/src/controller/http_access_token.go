package controller

import (
	"bookstoreapi/oauth/domain/access_token"
	at_services "bookstoreapi/oauth/services/access_token"
	"bookstoreapi/oauth/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByIdC(*gin.Context)
	CreateC(*gin.Context)
}

type accessTokenHandler struct {
	service at_services.Service
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

func (handler *accessTokenHandler) CreateC(c *gin.Context) {
	var atr access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&atr); err != nil {
		restErr := errors.NewBadRequestError("invalid fields")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := handler.service.Create(atr)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func NewHandler(service at_services.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}
