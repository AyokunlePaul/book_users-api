package users

import (
	"github.com/AyokunlePaul/book_users-api/domain/response"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (user *User) Validate() *response.BaseResponse {
	return nil
	//TODO Validate users email address pattern
}
