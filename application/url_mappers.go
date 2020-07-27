package application

import "github.com/AyokunlePaul/book_user-api/controllers"

func mapUrls() {
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:user_id", controllers.GetUser)
}
