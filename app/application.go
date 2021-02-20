package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-abdelrhman/go-phone-validator/utils/logger"
)

var (
	router=gin.Default()
)
func StartApplication()  {
	mapUrls()
	logger.Info("about to start the application")
	router.Run(":3000")
}