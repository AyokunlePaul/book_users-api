package application

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	log.Fatal(router.Run(":8080"))
}
