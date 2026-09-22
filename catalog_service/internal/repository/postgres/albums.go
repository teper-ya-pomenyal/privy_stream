package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

func (c *PostgresCatalog) GetAlbumByID(ctx context.Context, albumUUID uuid.UUID) (*domain.Album, error) {
	var album domain.Album
	err := c.conn.QueryRowContext(ctx, `
		SELECT
			album_id, artist_id, album_name, created_at
		FROM albums
		WHERE album_id = $1
		`, albumUUID,
	).Scan(&album.AlbumUUID, &album.ArtistUUID, &album.AlbumName, &album.CreatedAt)
	switch err {
	case sql.ErrNoRows:
		return nil, domain.ErrAlbumNotFound
	case nil:

		return &album, err
	default:
		return nil, err
	}

}

func (c *PostgresCatalog) GetAlbumTracks(ctx context.Context, albumUUID uuid.UUID) ([]domain.LightAlbumTrack, error) {
	// Трек принадлежит альбому либо через tracks.album_id (задаётся при AddTrack),
	// либо через albums_tracks (AddTracksToAlbum, там же позиция в трек-листе).
	// Учитываем оба источника: треки без позиции идут после пронумерованных.
	rows, err := c.conn.QueryContext(ctx, `
		SELECT t.track_id, t.track_name, t.explicit, t.duration_ms, COALESCE(at.position, 0)
		FROM tracks t
		LEFT JOIN albums_tracks at ON at.track_id = t.track_id AND at.album_id = $1
		WHERE t.album_id = $1 OR at.album_id IS NOT NULL
		ORDER BY at.position ASC NULLS LAST, t.created_at ASC
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

///////////////////////////////////////////////

func (c *PostgresCatalog) AddAlbum(ctx context.Context, album *domain.Album) error {
	_, err := c.conn.ExecContext(ctx, `
		INSERT INTO albums
			(album_id, artist_id, album_name, created_at)
		VALUES($1, $2, $3, $4)
		`, album.AlbumUUID, album.ArtistUUID, album.AlbumName, album.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				return domain.ErrAlbumAlreadyExists
			case "23503":
				return domain.ErrArtistNotFound
			}
		}
		return err
	}
	return nil
}
