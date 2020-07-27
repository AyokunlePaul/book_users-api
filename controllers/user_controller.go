package controllers

import (
	"fmt"
	"github.com/AyokunlePaul/book_user-api/domain/users"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func CreateUser(context *gin.Context) {
	user := users.User{}
	if bindError := context.ShouldBindJSON(&user); bindError != nil {
		for _, validationError := range bindError.(validator.ValidationErrors) {
			context.JSON(http.StatusBadRequest, gin.H{
				"successful": false,
				"message":    fmt.Sprintf("%s is missing", strings.Title(strings.ToLower(validationError.Field()))),
			})
			return
		}
	}
	rand.Seed(time.Now().UTC().UnixNano())
	userId := rand.Int63()
	user.Id = userId
	user.CreatedAt = time.Now()
	//TODO Save user to the database here!

	context.JSON(http.StatusCreated, gin.H{
		"successful": true,
		"message":    "User has been created successfully",
		"data":       user,
	})
}

func GetUser(context *gin.Context) {
	context.String(http.StatusNotImplemented, "Endpoint not implemented!")
}
