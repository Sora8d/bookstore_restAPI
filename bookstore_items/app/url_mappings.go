package app

import (
	"bookstoreapi/items/controllers"
	"net/http"
)

func Urlmaps() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	//router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
