package application

import "github.com/AyokunlePaul/book_users-api/controllers"

func mapUrls() {
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:user_id", controllers.GetUser)
}
