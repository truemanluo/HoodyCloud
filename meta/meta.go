package meta

import (
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Name        string
	Size        int64 // bytes
	ContentType string
	CreatedAt   time.Time
}

func GetFileInfo(filePath string) (*FileInfo, error) {
	// get file meta info
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	info := &FileInfo{
		Name:        fileInfo.Name(),
		Size:        fileInfo.Size(),
		ContentType: getContentType(filePath),
		CreatedAt:   fileInfo.ModTime(),
	}

	return info, nil
}

func getContentType(filePath string) string {
	// get file type by ext (MIME)
	ext := filepath.Ext(filePath)
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".pdf":
		return "application/pdf"
	default:
		return "application/octet-stream"
	}
}
