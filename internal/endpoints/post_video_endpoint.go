package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type PostVideoEndpoint struct {
}

func NewPostVideo() *PostVideoEndpoint {
	return &PostVideoEndpoint{}
}

func (endpoint *PostVideoEndpoint) PostVideoData(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	if !strings.HasSuffix(file.Filename, ".mp4") {
		context.JSON(http.StatusBadRequest, "file is not a mp4 file!")
		return
	}

	context.JSON(http.StatusOK, file.Filename)
}
