package services

import (
	"github.com/AyokunlePaul/book_users-api/domain/response"
	"github.com/AyokunlePaul/book_users-api/domain/users"
	"github.com/AyokunlePaul/book_users-api/utils"
)

var UserService userServiceInterface = &userService{}

type userService struct{}

type userServiceInterface interface {
	CreateUser(User users.User) (*users.User, *response.BaseResponse)
	GetUser(userId int64) (*users.User, *response.BaseResponse)
	GetAllUsers() ([]users.User, *response.BaseResponse)
	UpdateUser(User users.User, updateIsPartial bool) (*users.User, *response.BaseResponse)
	DeleteUser(userId int64) *response.BaseResponse
	GetUsersByStatus(status string) ([]users.User, *response.BaseResponse)
}

func (service *userService) CreateUser(User users.User) (*users.User, *response.BaseResponse) {
	if validationError := User.Validate(); validationError != nil {
		return nil, validationError
	}

	User.Status = utils.StatusEmailNotConfirmed
	createTime := utils.GetDBTime()
	User.CreatedAt = createTime
	User.UpdatedAt = createTime
	User.Password = utils.GetMD5(User.Password)

	if saveErrorResponse := User.Save(); saveErrorResponse != nil {
		return nil, saveErrorResponse
	}
	return &User, nil
}

func (service *userService) GetUser(userId int64) (*users.User, *response.BaseResponse) {
	if userId <= 0 {
		return nil, response.NewBadRequestError("invalid user id")
	}
	user := &users.User{
		Id: userId,
	}

	getErrorResponse := user.Get()
	if getErrorResponse != nil {
		return nil, getErrorResponse
	}

	return user, nil
}

func (service *userService) GetAllUsers() ([]users.User, *response.BaseResponse) {
	allUsers, getAllUsersError := users.GetAll()
	if getAllUsersError != nil {
		return nil, getAllUsersError
	}
	return allUsers, nil
}

func (service *userService) UpdateUser(User users.User, updateIsPartial bool) (*users.User, *response.BaseResponse) {
	currentUser, getErrorResponse := service.GetUser(User.Id)
	if getErrorResponse != nil {
		return nil, getErrorResponse
	}

	if updateIsPartial {
		currentUser.PartiallyUpdateUser(User)
	} else {
		if validationError := User.Validate(); validationError != nil {
			return nil, validationError
		}
		currentUser.FirstName = User.FirstName
		currentUser.LastName = User.LastName
		currentUser.Email = User.Email
		currentUser.UpdatedAt = utils.GetDBTime()
	}

	if updateErrorResponse := currentUser.Update(); updateErrorResponse != nil {
		return nil, updateErrorResponse
	}
	return currentUser, nil
}

func (service *userService) DeleteUser(userId int64) *response.BaseResponse {
	user, getUserError := service.GetUser(userId)
	if getUserError != nil {
		return getUserError
	}

	deleteError := user.Delete()
	if deleteError != nil {
		return deleteError
	}

	return nil
}

func (service *userService) GetUsersByStatus(status string) ([]users.User, *response.BaseResponse) {
	return users.GetUsersByStatus(status)
}
