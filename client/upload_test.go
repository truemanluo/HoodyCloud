package client

import "testing"

func TestXxx(t *testing.T) {
	if err := UploadFile("http://127.0.0.1:8080/file/upload", "upload.go", "file"); err != nil {
		t.Fatalf("upload returned err: %v", err.Error())
	}
}
