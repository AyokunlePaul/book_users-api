package application

import (
	"github.com/AyokunlePaul/book_users-api/logger"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var (
	router = gin.New()
)

func init() {
	zapLogger := logger.GetLogger()
	router.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(zapLogger, true))
}

func StartApplication() {
	mapUrls()
	log.Fatal(router.Run(":8080"))
}
