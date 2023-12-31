package main

import (
	"fmt"
	"net/http"

	"github.com/truemanluo/HoodyCloud/handler"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("failed to start server: %s", err.Error())
	}
}
