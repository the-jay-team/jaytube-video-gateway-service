package message_queue_provider

type MessageQueueProvider interface {
	AddToQueue(topic *string, message []byte) error
}
