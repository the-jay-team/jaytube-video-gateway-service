package main

import (
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/client"
	"github.com/the-jay-team/jaytube-upload-service/internal/endpoints"
	"github.com/the-jay-team/jaytube-upload-service/test/message_queue_provider"
	"log"
)

func main() {
	server := gin.Default()

	mockedMessageQueue := message_queue_provider.NewMockedMessageQueueProvider()
	mockedIrisClient := client.NewMockedContentInformationServiceClient()
	videoPostEndpoint := endpoints.NewPostVideo(mockedMessageQueue, mockedIrisClient)
	videoDeleteEndpoint := endpoints.NewDeleteVideoEndpoint(mockedIrisClient)

	server.POST("/video/:id", videoPostEndpoint.PostVideo)
	server.DELETE("/video/:id", videoDeleteEndpoint.DeleteVideo)

	ginStartError := server.Run(":8080")
	if ginStartError != nil {
		log.Fatalf("Could not Startup gin: %s", ginStartError)
	}
}
