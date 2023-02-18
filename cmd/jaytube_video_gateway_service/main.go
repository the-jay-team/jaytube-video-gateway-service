package main

import (
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/client"
	"github.com/the-jay-team/jaytube-upload-service/internal/configs"
	"github.com/the-jay-team/jaytube-upload-service/internal/endpoints"
	"github.com/the-jay-team/jaytube-upload-service/internal/message_queue_provider"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
)

func main() {
	config := configs.GetEnvironmentConfig()
	server := gin.Default()
	hostname, _ := os.Hostname()
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.Target,
		"client.id":         hostname,
		"acks":              "all"})
	if err != nil {
		println(err.Error())
	}
	kafkaMessageQueue := message_queue_provider.NewKafkaMessageQueueProvider(kafkaProducer)
	irisClient := client.NewContentInformationServiceClient(config.IrisClient.Target)
	videoPostEndpoint := endpoints.NewPostVideo(kafkaMessageQueue, irisClient)
	videoDeleteEndpoint := endpoints.NewDeleteVideoEndpoint(irisClient)

	server.POST("/video/:id", videoPostEndpoint.PostVideo)
	server.DELETE("/video/:id", videoDeleteEndpoint.DeleteVideo)

	ginStartError := server.Run(":8080")
	if ginStartError != nil {
		log.Fatalf("Could not Startup gin: %s", ginStartError)
	}
}
