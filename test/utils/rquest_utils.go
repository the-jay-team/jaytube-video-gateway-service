package utils

import (
	"bytes"
	"mime/multipart"
)

func CreateMultipartWriterWithFile(fieldName string, fileName string) (*multipart.Writer, *bytes.Buffer) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, _ = writer.CreateFormFile(fieldName, fileName)
	_ = writer.Close()
	return writer, body
}
