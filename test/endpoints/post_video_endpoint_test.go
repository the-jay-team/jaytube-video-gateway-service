package endpoints

import (
	"bytes"
	"github.com/go-playground/assert/v2"
	"github.com/the-jay-team/jaytube-upload-service/internal/endpoints"
	"github.com/the-jay-team/jaytube-upload-service/test"
	"mime/multipart"
	"net/http"
	"testing"
)

func TestPostWrongFileFormatReturnsBadRequest(t *testing.T) {
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodPost, "/video", nil)

	testEndpoint := endpoints.NewPostVideo()
	testEndpoint.PostVideoData(context)

	assert.Equal(t, http.StatusBadRequest, record.Code)
}

func TestPostVideoReturnsOk(t *testing.T) {
	record, context := test.GinTestSetup()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, _ = writer.CreateFormFile("video", "test.mp4")
	_ = writer.Close()

	context.Request, _ = http.NewRequest(http.MethodPost, "/video", bytes.NewReader(body.Bytes()))
	context.Request.Header.Set("Content-Type", "multipart/form-data")
	
	testEndpoint := endpoints.NewPostVideo()
	testEndpoint.PostVideoData(context)

	assert.Equal(t, http.StatusBadRequest, record.Code)
}
