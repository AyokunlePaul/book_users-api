package services

import (
	"github.com/AyokunlePaul/book_users-api/domain/response"
	"github.com/AyokunlePaul/book_users-api/domain/users"
)

func CreateUser(User users.User) (*users.User, *response.BaseResponse) {
	if saveErrorResponse := User.SaveUser(); saveErrorResponse != nil {
		return nil, saveErrorResponse
	}
	return &User, nil
}

func GetUser(userId int64) (*users.User, *response.BaseResponse) {
	if userId <= 0 {
		return nil, response.NewBadRequestError("Invalid user id")
	}
	user := &users.User{
		Id: userId,
	}
	getErrorResponse := user.GetUser()
	if getErrorResponse != nil {
		return nil, getErrorResponse
	}

	return user , nil
}
