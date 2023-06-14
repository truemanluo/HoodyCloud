package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	// receive form
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("failed to get data: %s", err.Error())
		return
	}
	defer file.Close()

	// store file locally
	newFile, err := os.Create("/tmp/" + header.Filename)
	if err != nil {
		log.Printf("failed to create local file: %s", err.Error())
		return
	}
	defer newFile.Close()

	// copy form file to local file
	_, err = io.Copy(newFile, file)
	if err != nil {
		log.Printf("failed to copy local file: %s", err.Error())
		return
	}
}

// download file by url
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// get file url
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing URL parameter", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to download file: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	_, filename := filepath.Split(url)

	// set response headers
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	// write file content to response body
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to write file to response: %v", err), http.StatusInternalServerError)
		return
	}
}
