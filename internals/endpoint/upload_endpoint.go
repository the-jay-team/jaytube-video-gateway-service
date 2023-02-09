package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"jaytube-upload-service/pkg/data"
	"log"
	"net/http"
	"os"
)

func UploadVideo(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	err := c.SaveUploadedFile(file, os.Getenv("FILE_DESTINATION")+"/"+file.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
}

func CreateVideo(c *gin.Context) {
	var videoData data.VideoDataPayload
	if c.ShouldBindBodyWith(&videoData, binding.JSON) != nil {
		c.JSON(http.StatusBadRequest, "Wrong JSON request body")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"json": videoData,
	})
}
