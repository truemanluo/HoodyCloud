package client

import "testing"

func TestXxx(t *testing.T) {
	UploadFile("http://127.0.0.1:8080/file/upload", "upload.go", "file")
}
