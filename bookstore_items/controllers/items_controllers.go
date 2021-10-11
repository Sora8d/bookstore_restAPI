package controllers

import (
	"bookstoreapi/items/domain/items"
	"bookstoreapi/items/services"
	"bookstoreapi/items/utils/http_utils"
	"net/http"

	"github.com/Sora8d/bookstore_oauth-go/oauth"
)

var ItemsController itemsControllerInterface = &itemsController{}

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct {
}

func (it *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: Return error to the caller
		http_utils.RespondJson(w, err.Status(), err)
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO RESTURN ERROR json to the user
		http_utils.RespondJson(w, err.Status(), err)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
	//TODO: Return created item with HTTP
}

func (it *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	return
}
