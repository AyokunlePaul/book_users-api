package users

import (
	"github.com/AyokunlePaul/book_users-api/domain/response"
	"github.com/AyokunlePaul/book_users-api/domain/users"
	"github.com/AyokunlePaul/book_users-api/services"
	"github.com/AyokunlePaul/book_users-api/utils"
	"github.com/AyokunlePaul/book_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func Create(context *gin.Context) {
	user := users.User{}
	if bindError := context.ShouldBindJSON(&user); bindError != nil {
		validationError := errors.ValidationError(bindError)
		context.JSON(validationError.Status, validationError)
		return
	}

	result, serviceError := services.UserService.CreateUser(user)
	if serviceError != nil {
		context.JSON(serviceError.Status, serviceError)
		return
	}

	successMessage := "user successfully created"
	context.JSON(http.StatusCreated, response.NewCreateResponse(successMessage, result.Marshal()))
}

func Get(context *gin.Context) {
	userId, parseError := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if parseError != nil {
		context.JSON(http.StatusBadRequest, response.NewBadRequestError("invalid id"))
		return
	}

	user, getUserBaseResponse := services.UserService.GetUser(userId)
	if getUserBaseResponse != nil {
		context.JSON(getUserBaseResponse.Status, getUserBaseResponse)
		return
	}

	context.JSON(http.StatusOK, response.NewOkResponse("user details fetched successfully", user.Marshal()))
}

func GetAll(context *gin.Context) {
	allUsers, getAllUsersError := services.UserService.GetAllUsers()
	if getAllUsersError != nil {
		context.JSON(getAllUsersError.Status, getAllUsersError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("all users fetched successfully", users.Users(allUsers).Marshal()))
}

func Update(context *gin.Context) {
	userId, parseError := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if parseError != nil {
		context.JSON(http.StatusBadRequest, response.NewBadRequestError("invalid id"))
		return
	}

	user := users.User{}
	if bindError := context.ShouldBindJSON(&user); bindError != nil {
		validationError := errors.ValidationError(bindError)
		context.JSON(validationError.Status, validationError)
		return
	}

	updateIsPartial := context.Request.Method == http.MethodPatch

	user.Id = userId
	updateUser, updateUserResponse := services.UserService.UpdateUser(user, updateIsPartial)
	if updateUserResponse != nil {
		context.JSON(updateUserResponse.Status, updateUserResponse)
		return
	}

	message := "user details updated successfully"
	context.JSON(http.StatusOK, response.NewOkResponse(message, updateUser.Marshal()))
}

func Delete(context *gin.Context) {
	userId, parseError := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if parseError != nil {
		context.JSON(http.StatusBadRequest, response.NewBadRequestError("invalid id"))
		return
	}
	deleteUserBaseResponse := services.UserService.DeleteUser(userId)
	if deleteUserBaseResponse != nil {
		context.JSON(deleteUserBaseResponse.Status, deleteUserBaseResponse)
		return
	}

	context.JSON(http.StatusOK, response.NewOkResponse("user successfully deleted", nil))
}

func GetByStatus(context *gin.Context) {
	status := strings.TrimSpace(context.Query("status"))
	if status == "" || utils.IsNotValidStatus(status) {
		message := "invalid status"
		context.JSON(http.StatusBadRequest, response.NewBadRequestError(message))
		return
	}
	usersWithStatus, getUsersWithStatusError := services.UserService.GetUsersByStatus(status)
	if getUsersWithStatusError != nil {
		context.JSON(getUsersWithStatusError.Status, getUsersWithStatusError)
		return
	}

	message := "users fetched"
	context.JSON(http.StatusOK, response.NewOkResponse(message, users.Users(usersWithStatus).Marshal()))
}
