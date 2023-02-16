package main

import (
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-upload-service/internal/endpoints"
	"log"
)

func main() {
	server := gin.Default()

	server.POST("/video", endpoints.NewPostVideo().PostVideoData)

	ginStartError := server.Run(":8080")
	if ginStartError != nil {
		log.Fatalf("Could not Startup gin: %s", ginStartError)
	}
}
