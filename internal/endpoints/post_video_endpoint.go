package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/client"
	"github.com/the-jay-team/jaytube-upload-service/internal/message_queue_provider"
	"net/http"
	"strings"
)

type PostVideoEndpoint struct {
	queue message_queue_provider.MessageQueueProvider
	iris  client.IrisClient
}

func NewPostVideo(queue message_queue_provider.MessageQueueProvider, irisClient client.IrisClient) *PostVideoEndpoint {
	return &PostVideoEndpoint{queue, irisClient}
}

func (endpoint *PostVideoEndpoint) PostVideo(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, "Missing video id!")
		return
	}
	idExists, irisErr := endpoint.iris.VideoExists(id)
	if irisErr != nil {
		context.JSON(http.StatusInternalServerError, fmt.Sprintf("could not get video data from iris: %s",
			irisErr.Error()))
		return
	}
	if !idExists {
		context.JSON(http.StatusNotFound, "Video id does not exist!")
		return
	}

	file, fileErr := context.FormFile("video")
	if fileErr != nil {
		context.JSON(http.StatusInternalServerError, fileErr.Error())
		return
	}
	if !strings.HasSuffix(file.Filename, ".mp4") {
		context.JSON(http.StatusBadRequest, "file is not a mp4 file!")
		return
	}

	_ = context.SaveUploadedFile(file, fmt.Sprintf("uploads/%s.mp4", id))
	topic := "video"
	providerErr := endpoint.queue.AddToQueue(&topic, []byte(id))
	if providerErr != nil {
		context.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to queue video: %s", providerErr.Error()))
		return
	}

	context.JSON(http.StatusOK, "")
}
