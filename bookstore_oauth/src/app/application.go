package app

import (
	"bookstoreapi/oauth/clients/cassandra"
	"bookstoreapi/oauth/controller"
	"bookstoreapi/oauth/domain/access_token"
	"bookstoreapi/oauth/repository/db"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	atService := access_token.NewService(db.NewRepository())
	atHandler := controller.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByIdC)
	router.Run("localhost:8080")
}
