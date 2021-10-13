package controllers

import (
	"bookstoreapi/items/domain/items"
	"bookstoreapi/items/domain/queries"
	"bookstoreapi/items/localerrors"
	"bookstoreapi/items/services"
	"bookstoreapi/items/utils/http_utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Sora8d/bookstore_oauth-go/oauth"
	"github.com/Sora8d/bookstore_utils-go/rest_errors"
	"github.com/gorilla/mux"
)

var ItemsController itemsControllerInterface = &itemsController{}

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}

type itemsController struct {
}

func (it *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: Return error to the caller
		http_utils.RespondJson(w, err.Status(), err)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := localerrors.NewUnauthorizedError("invalid request body")
		http_utils.RespondJson(w, respErr.Status, respErr)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestErr("invalid request body")
		http_utils.RespondJson(w, respErr.Status(), respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestErr("invalid json body")
		http_utils.RespondJson(w, respErr.Status(), respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, respErr := services.ItemsService.Create(itemRequest)
	if respErr != nil {
		//TODO RESTURN ERROR json to the user
		http_utils.RespondJson(w, respErr.Status(), respErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
	//TODO: Return created item with HTTP
}

func (it *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.RespondJson(w, err.Status(), err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}

func (it *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestErr("invalid json body")
		http_utils.RespondJson(w, respErr.Status(), respErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		respErr := rest_errors.NewBadRequestErr("invalid json body")
		http_utils.RespondJson(w, respErr.Status(), respErr)
		return
	}

	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.RespondJson(w, searchErr.Status(), searchErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, items)
}
