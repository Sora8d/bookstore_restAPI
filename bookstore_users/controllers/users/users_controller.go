package users

import (
	"bookstoreapi/users/domain/users"
	"bookstoreapi/users/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//This is going to handle controllers

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	//First way
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//TODO: Handle error
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation error
		return
	}
	fmt.Println(result)
	c.String(http.StatusNotImplemented, "implement me!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
