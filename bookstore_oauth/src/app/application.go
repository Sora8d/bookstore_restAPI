package app

import (
	"bookstoreapi/oauth/controller"
	"bookstoreapi/oauth/repository/db"
	"bookstoreapi/oauth/repository/rest"
	at_services "bookstoreapi/oauth/services/access_token"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	atService := at_services.NewService(db.NewRepository(), rest.NewRepository())
	atHandler := controller.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByIdC)
	router.POST("/oauth/access_token", atHandler.CreateC)
	router.Run("localhost:8081")
}
