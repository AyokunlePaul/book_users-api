package application

import (
	"github.com/AyokunlePaul/book_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/users", users.GetAll)
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/search/users", users.GetByStatus)
}
