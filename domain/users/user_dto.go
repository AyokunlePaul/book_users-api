package users

import (
	"encoding/json"
	"github.com/AyokunlePaul/book_users-api/domain/response"
	"github.com/AyokunlePaul/book_users-api/utils"
	"strings"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (user *User) Validate() *response.BaseResponse {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return response.NewBadRequestError("email address is required")
	}
	if !utils.IsValidEmail(user.Email) {
		return response.NewBadRequestError("email is not valid")
	}
	if user.Password == "" {
		return response.NewBadRequestError("password is required")
	}
	return nil
}

func (user *User) PartiallyUpdateUser(from User) {
	if from.FirstName != "" {
		user.FirstName = from.FirstName
	}
	if from.LastName != "" {
		user.LastName = from.LastName
	}
	if from.Email != "" || utils.IsValidEmail(from.Email) {
		user.Email = from.Email
	}
}

func (user *User) Marshal() interface{} {
	var publicUser PublicUser
	userJson, _ := json.Marshal(user)
	_ = json.Unmarshal(userJson, &publicUser)

	return publicUser
}

func (users Users) Marshal() interface{} {
	result := make([]interface{}, len(users))
	for index, value := range users {
		result[index] = value.Marshal()
	}

	return result
}
