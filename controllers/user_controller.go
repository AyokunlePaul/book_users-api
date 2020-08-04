package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/AyokunlePaul/book_users-api/domain/response"
	"github.com/AyokunlePaul/book_users-api/domain/users"
	"github.com/AyokunlePaul/book_users-api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CreateUser(context *gin.Context) {
	user := users.User{}
	if bindError := context.ShouldBindJSON(&user); bindError != nil {
		switch errorType := bindError.(type) {
		case *json.UnmarshalTypeError:
			message := fmt.Sprintf("%s is invalid", errorType.Field)
			context.JSON(http.StatusBadRequest, response.NewBadRequestError(message))
			return
		case validator.ValidationErrors:
			for _, validationError := range errorType {
				validationMessage := fmt.Sprintf("%s is a required field", strings.Title(strings.ToLower(validationError.Field())))
				context.JSON(http.StatusBadRequest, response.NewBadRequestError(validationMessage))
				return
			}
		default:
			context.JSON(http.StatusBadRequest, response.NewBadRequestError("Cannot process fields. Please check and try again"))
			return
		}
	}

	user.CreatedAt = time.Now()

	result, serviceError := services.CreateUser(user)
	if serviceError != nil {
		context.JSON(serviceError.Status, serviceError)
		return
	}

	successMessage := "User successfully created"
	context.JSON(http.StatusCreated, response.NewCreateResponse(successMessage, result))
}

func GetUser(context *gin.Context) {
	userId, parseError := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if parseError != nil {
		context.JSON(http.StatusBadRequest, response.NewBadRequestError("The supplied ID is invalid"))
		return
	}

	user, getUserBaseResponse := services.GetUser(userId)
	if getUserBaseResponse != nil {
		context.JSON(getUserBaseResponse.Status, getUserBaseResponse)
		return
	}

	context.JSON(http.StatusOK, response.NewOkResponse("User details fetched successfully", user))
}
