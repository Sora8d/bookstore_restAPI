package users

import (
	"bookstoreapi/users/domain/users"
	"bookstoreapi/users/services"
	"bookstoreapi/users/utils/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	/*We can also use c.ShouldBindJSON(&user), that replaces everythin from
	Readall() to Unmarshal()*/

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}
	/* albeit this works, it breaks the way structures are arranged, so the video creates and uses a function from the service package
	reqUser := users.User{Id: userId}
	getErr := reqUser.Get()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	*/

	reqUser, reqErr := services.GetUser(userId)
	if reqErr != nil {
		c.JSON(reqErr.Status, reqErr)
		return
	}

	/* This is not needed, c.JSON takes care of transforming the struct to JSON, but in your own implementation of a router it will be useful
	retUserJSON, jsonErr := json.Marshal(reqUser)
	if jsonErr != nil {
		//TODO: implement marshal error (Does this need an implementation?)
		return
	}
	*/
	c.JSON(http.StatusOK, reqUser)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
