package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"jaytube-upload-service/internals/endpoint"
	"jaytube-upload-service/pkg/env_var_util"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	err = env_var_util.CheckEnvVars([]string{"FILE_DESTINATION"})
	if err != nil {
		return
	}

	r := gin.Default()

	r.POST("/upload", endpoint.UploadVideo)

	err = r.Run(":8083")
	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
