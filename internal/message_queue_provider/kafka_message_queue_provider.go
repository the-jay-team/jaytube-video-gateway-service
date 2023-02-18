package message_queue_provider

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

type KafkaMessageQueueProvider struct {
	producer *kafka.Producer
}

func NewKafkaMessageQueueProvider(producer *kafka.Producer) *KafkaMessageQueueProvider {
	return &KafkaMessageQueueProvider{producer}
}

func (provider *KafkaMessageQueueProvider) AddToQueue(topic *string, message []byte) error {
	deliveryChan := make(chan kafka.Event, 10000)
	err := provider.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)

	if err != nil {
		return err
	}
	return nil
}
