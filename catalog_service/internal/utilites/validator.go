package utilites

import (
	"strings"

	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

func ValidateTrackName(trackName string) (string, error) {
	cleanTN := strings.TrimSpace(trackName)
	if cleanTN == "" {
		return "", domain.ErrInvalidTrackName
	}
	return cleanTN, nil
}

func ValidateArtistName(artistName string) (string, error) {
	cleanAN := strings.TrimSpace(artistName)
	if cleanAN == "" {
		return "", domain.ErrInvalidArtistName
	}
	return cleanAN, nil
}

func ValidateAlbumName(albumName string) (string, error) {
	cleanAlN := strings.TrimSpace(albumName)
	if cleanAlN == "" {
		return "", domain.ErrInvalidAlbumName
	}
	return cleanAlN, nil
}
