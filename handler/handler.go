package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	// 接收文件流并存储到本地
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Printf("failed to get data: %s", err.Error())
		return
	}
	defer file.Close()

	// 创建本地文件
	newFile, err := os.Create("/tmp/" + header.Filename)
	if err != nil {
		fmt.Printf("failed to create local file: %s", err.Error())
		return
	}
	defer newFile.Close()

	// 将上传文件流拷贝到本地文件
	_, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Printf("failed to copy local file: %s", err.Error())
		return
	}
}