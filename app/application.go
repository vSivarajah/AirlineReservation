package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {
	MapUrls()
	log.Printf("Starting application server...")
	router.Run(":8081")
}
