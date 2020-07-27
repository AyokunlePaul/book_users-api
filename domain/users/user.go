package users

import "time"

type User struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"first_name" binding:"required" tag:"first_name"`
	LastName  string    `json:"last_name" binding:"required" tag:"last_name"`
	Email     string    `json:"email" binding:"required" tag:"email"`
	CreatedAt time.Time `json:"created_at"`
}
