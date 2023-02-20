package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/client"
	"net/http"
	"os"
)

type DeleteVideoEndpoint struct {
	iris client.IrisClient
}

func NewDeleteVideoEndpoint(irisClient client.IrisClient) *DeleteVideoEndpoint {
	return &DeleteVideoEndpoint{irisClient}
}

func (endpoint *DeleteVideoEndpoint) DeleteVideo(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, "Missing video id!")
		return
	}
	idExists, irisErr := endpoint.iris.VideoExists(id)
	if irisErr != nil {
		context.JSON(http.StatusInternalServerError, fmt.Sprintf("Could not get video data from iris: %s",
			irisErr.Error()))
		return
	}
	if !idExists {
		context.JSON(http.StatusNotFound, "Video id does not exist!")
		return
	}

	removeErr := os.RemoveAll(fmt.Sprintf("storage/%s", id))
	if removeErr != nil {
		context.JSON(http.StatusInternalServerError, fmt.Sprintf("could not delete video file: %s",
			removeErr.Error()))
		return
	}
	_, irisDelErr := endpoint.iris.DeleteVideoData(id)
	if irisDelErr != nil {
		context.JSON(http.StatusInternalServerError, fmt.Sprintf("error occured deleting video data: %s",
			irisDelErr.Error()))
		return
	}

	context.JSON(http.StatusOK, "")
}
