package usecase

import (
	"os"
	"path/filepath"
	"time"

	"github.com/teper-ya-pomenyal/privy_stream/streaming_service/internal/client/grpc"
	"github.com/teper-ya-pomenyal/privy_stream/streaming_service/internal/domain"
	"github.com/teper-ya-pomenyal/privy_stream/streaming_service/internal/storage"
)

type StreamTrackResponse struct {
	File       *os.File
	FileName   string
	ModTime    time.Time
	DurationMS int32
}

type Streamer struct {
	trackStoragePath string
}

func NewStreamer(storagePath string) *Streamer {
	return &Streamer{trackStoragePath: storagePath}
}

// isAdult — исполнилось ли 18 полных лет на момент now.
func isAdult(birthDate, now time.Time) bool {
	return !now.Before(birthDate.AddDate(18, 0, 0))
}

func (s *Streamer) StreamTrack(tl *grpc.TrackLocation, bd time.Time) (*StreamTrackResponse, error) {
	if tl.Explicit && !isAdult(bd, time.Now()) {
		return nil, domain.ErrExplicitContentBlocked
	}
	trackPath := filepath.Join(s.trackStoragePath, tl.Path)
	fr, err := storage.ReadFile(trackPath)
	if err != nil {
		return nil, err
	}
	return &StreamTrackResponse{
		File:       fr.File,
		FileName:   fr.File.Name(),
		ModTime:    fr.ModTime,
		DurationMS: tl.DurationMS,
	}, nil
}
