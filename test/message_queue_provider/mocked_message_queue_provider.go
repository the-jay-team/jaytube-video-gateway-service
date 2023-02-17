package message_queue_provider

type MockedMessageQueueProvider struct {
}

func NewMockedMessageQueueProvider() *MockedMessageQueueProvider {
	return &MockedMessageQueueProvider{}
}

func (provider *MockedMessageQueueProvider) AddToQueue(topic *string, message []byte) error {
	return nil
}
