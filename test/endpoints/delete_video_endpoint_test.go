package endpoints

import (
	"github.com/go-playground/assert/v2"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/client"
	"github.com/the-jay-team/jaytube-upload-service/internal/endpoints"
	"github.com/the-jay-team/jaytube-upload-service/test/utils"
	"net/http"
	"testing"
)

func TestDeleteVideoReturnsOK(t *testing.T) {
	record, context := utils.GinTestSetup()

	context.Request, _ = http.NewRequest(http.MethodDelete, "video/:id", nil)
	context.AddParam("id", "1")

	testEndpoint := endpoints.NewDeleteVideoEndpoint(client.NewMockedContentInformationServiceClient())
	testEndpoint.DeleteVideo(context)

	assert.Equal(t, http.StatusOK, record.Code)
}

func TestMissingIdReturnBadRequest(t *testing.T) {
	record, context := utils.GinTestSetup()

	context.Request, _ = http.NewRequest(http.MethodDelete, "/video/:id", nil)

	testEndpoint := endpoints.NewDeleteVideoEndpoint(client.NewMockedContentInformationServiceClient())
	testEndpoint.DeleteVideo(context)

	assert.Equal(t, http.StatusBadRequest, record.Code)
}

func TestVideoDoesNotExistReturnNotFound(t *testing.T) {
	record, context := utils.GinTestSetup()

	context.Request, _ = http.NewRequest(http.MethodPost, "/video/:id", nil)
	context.AddParam("id", "awdawfdf")

	testEndpoint := endpoints.NewDeleteVideoEndpoint(client.NewMockedContentInformationServiceClient())
	testEndpoint.DeleteVideo(context)

	assert.Equal(t, http.StatusNotFound, record.Code)
}
