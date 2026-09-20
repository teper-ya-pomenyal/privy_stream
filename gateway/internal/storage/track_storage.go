package storage

import (
	"io"
	"os"
	"path/filepath"
)

type TrackStorage struct {
	basePath string
}

func NewTrackStorage(basePath string) *TrackStorage {
	return &TrackStorage{basePath: basePath}
}

func (s *TrackStorage) AddTrackFile(path string, r io.Reader) (int64, error) {
	fullPath := filepath.Join(s.basePath, path)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
		return 0, err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return io.Copy(file, r)
}
