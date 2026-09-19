package storage

import (
	"errors"
	"io/fs"
	"os"
	"time"

	"github.com/teper-ya-pomenyal/privy_stream/streaming_service/internal/domain"
)

type FileReaderResponse struct {
	File    *os.File
	ModTime time.Time
}

func ReadFile(filePath string) (*FileReaderResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		switch {
		case errors.Is(err, fs.ErrNotExist):
			return nil, domain.ErrFileNotExists
		case errors.Is(err, fs.ErrPermission):
			return nil, domain.ErrNotPermission
		default:
			return nil, err
		}
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	md := fi.ModTime()
	res := &FileReaderResponse{File: file, ModTime: md}
	return res, nil
}
