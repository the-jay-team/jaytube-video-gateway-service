package main

import (
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-upload-service/internal/endpoints"
	"github.com/the-jay-team/jaytube-upload-service/internal/message_queue_provider"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
)

func main() {
	server := gin.Default()
	hostname, _ := os.Hostname()
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "dev.alaust.azubi.server.lan:3307",
		"client.id":         hostname,
		"acks":              "all"})
	if err != nil {
		println(err.Error())
	}
	kafkaMessageQueue := message_queue_provider.NewKafkaMessageQueueProvider(kafkaProducer)
	videoPostEndpoint := endpoints.NewPostVideo(kafkaMessageQueue)

	server.POST("/video/:id", videoPostEndpoint.PostVideo)

	ginStartError := server.Run(":8080")
	if ginStartError != nil {
		log.Fatalf("Could not Startup gin: %s", ginStartError)
	}
}
