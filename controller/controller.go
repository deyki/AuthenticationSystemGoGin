package controller

import (
	"net/http"

	"github.com/AuthSystemJWT/deyki/v2/database"
	"github.com/AuthSystemJWT/deyki/v2/middleware"
	"github.com/AuthSystemJWT/deyki/v2/service"
	"github.com/gin-gonic/gin"
)

func signUpHandler(c *gin.Context) {

	var user database.User

	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read json body"})
		return
	}

	signUpResponse, errorMessage := service.SignUp(&user)
	if errorMessage != nil {
		c.IndentedJSON(errorMessage.HttpStatus, gin.H{"error": errorMessage})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"response": signUpResponse})
}


func signInHandler(c *gin.Context) {

	var requestBodyUser database.User

	if c.BindJSON(&requestBodyUser) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	singInResponse, errorMessage := service.SignIn(&requestBodyUser)
	if errorMessage != nil {
		c.IndentedJSON(errorMessage.HttpStatus, gin.H{"error": errorMessage})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", singInResponse.JWToken, 3600 * 24 * 30, "", "", false, true)

	c.IndentedJSON(http.StatusOK, gin.H{"Authenticated": true})
}


func HomeHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Home page!"})
}

func GinRouter() {
	router := gin.Default()
	router.POST("/sign-up", signUpHandler)
	router.POST("/sign-in", signInHandler)
	router.GET("/home", middleware.RequireAuth, HomeHandler)
	router.Run()
}