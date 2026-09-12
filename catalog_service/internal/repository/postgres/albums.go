package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

func (c *PostgresCatalog) GetAlbumByID(ctx context.Context, albumUUID uuid.UUID) (*domain.Album, error) {
	var album domain.Album
	err := c.conn.QueryRowContext(ctx, `
		SELECT
			al.album_id, al.artist_id, al.album_name,
			ar.artist_name, al.created_at
		FROM albums al
		JOIN artists ar ON ar.artist_id = al.artist_id
		WHERE al.album_id = $1
		`, albumUUID,
	).Scan(&album.AlbumID, &album.ArtistID, &album.AlbumName, &album.ArtistName, &album.CreatedAt)
	switch err {
	case sql.ErrNoRows:
		return nil, domain.ErrAlbumNotFound
	case nil:
		album.Tracks, err = c.GetAlbumTracks(ctx, albumUUID)
		if err != nil {
			return nil, err
		}
		return &album, err
	default:
		return nil, err
	}

}

func (c *PostgresCatalog) GetAlbumTracks(ctx context.Context, albumUUID uuid.UUID) ([]domain.LightAlbumTrack, error) {
	rows, err := c.conn.QueryContext(ctx, `
		SELECT at.track_id, t.track_name, t.explicit, t.duration_ms, at.position
		FROM albums_tracks at
		JOIN tracks t ON t.track_id = at.track_id
		WHERE at.album_id = $1
		ORDER BY at.position ASC
		`,
		albumUUID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tracks := []domain.LightAlbumTrack{}
	for rows.Next() {
		var t domain.LightAlbumTrack
		if err := rows.Scan(&t.TrackID, &t.TrackName, &t.Explicit, &t.DurationMS, &t.Position); err != nil {
			return nil, err
		}
		tracks = append(tracks, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tracks, nil
}
