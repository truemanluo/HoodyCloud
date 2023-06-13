package client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// generate by chatgpt
func UploadFile(url, filePath, formName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("failed to open file %s: %v\n", filePath, err)
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(formName, filepath.Base(filePath))
	if err != nil {
		log.Printf("failed to create form file: %v\n", err)
		return err
	}

	if _, err := io.Copy(part, file); err != nil {
		log.Printf("failed to copy file data to form file: %v\n", err)
		return err
	}

	if err := writer.Close(); err != nil {
		log.Printf("failed to close multipart writer: %v\n", err)
		return err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Printf("failed to create HTTP request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("failed to upload file: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("failed to upload file, HTTP status code: %d\n", resp.StatusCode)
		return fmt.Errorf("failed to upload file, HTTP status code: %d", resp.StatusCode)
	}

	return nil
}
