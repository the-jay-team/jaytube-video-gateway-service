package endpoints

import (
	"github.com/go-playground/assert/v2"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/client"
	"github.com/the-jay-team/jaytube-upload-service/internal/endpoints"
	"github.com/the-jay-team/jaytube-upload-service/test/message_queue_provider"
	"github.com/the-jay-team/jaytube-upload-service/test/utils"
	"net/http"
	"testing"
)

func TestPostWrongFileFormatReturnsBadRequest(t *testing.T) {
	record, context := utils.GinTestSetup()

	writer, body := utils.CreateMultipartWriterWithFile("video", "test.png")

	context.Request, _ = http.NewRequest(http.MethodPost, "/video/:id", body)
	context.AddParam("id", "1")
	context.Request.Header.Set("Content-Type", writer.FormDataContentType())

	testEndpoint := endpoints.NewPostVideo(message_queue_provider.NewMockedMessageQueueProvider(),
		client.NewMockedContentInformationServiceClient())
	testEndpoint.PostVideo(context)

	assert.Equal(t, http.StatusBadRequest, record.Code)
}

func TestPostVideoReturnsOk(t *testing.T) {
	record, context := utils.GinTestSetup()

	writer, body := utils.CreateMultipartWriterWithFile("video", "test.mp4")

	context.Request, _ = http.NewRequest(http.MethodPost, "/video/:id", body)
	context.AddParam("id", "1")
	context.Request.Header.Set("Content-Type", writer.FormDataContentType())

	testEndpoint := endpoints.NewPostVideo(message_queue_provider.NewMockedMessageQueueProvider(),
		client.NewMockedContentInformationServiceClient())
	testEndpoint.PostVideo(context)

	assert.Equal(t, http.StatusOK, record.Code)
}

func TestMissingVideoIdReturnsBadRequest(t *testing.T) {
	record, context := utils.GinTestSetup()

	writer, body := utils.CreateMultipartWriterWithFile("video", "test.mp4")

	context.Request, _ = http.NewRequest(http.MethodPost, "/video/:id", body)
	context.Request.Header.Set("Content-Type", writer.FormDataContentType())

	testEndpoint := endpoints.NewPostVideo(message_queue_provider.NewMockedMessageQueueProvider(),
		client.NewMockedContentInformationServiceClient())
	testEndpoint.PostVideo(context)

	assert.Equal(t, http.StatusBadRequest, record.Code)
}

func TestVideoIdDoesNotExistReturnsNotFound(t *testing.T) {
	record, context := utils.GinTestSetup()

	writer, body := utils.CreateMultipartWriterWithFile("video", "test.mp4")

	context.Request, _ = http.NewRequest(http.MethodPost, "/video/:id", body)
	context.AddParam("id", "awdawfdf")
	context.Request.Header.Set("Content-Type", writer.FormDataContentType())

	testEndpoint := endpoints.NewPostVideo(message_queue_provider.NewMockedMessageQueueProvider(),
		client.NewMockedContentInformationServiceClient())
	testEndpoint.PostVideo(context)

	assert.Equal(t, http.StatusNotFound, record.Code)
}
