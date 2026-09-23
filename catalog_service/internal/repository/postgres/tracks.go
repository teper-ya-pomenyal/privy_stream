package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

func (c *PostgresCatalog) GetTrackByID(ctx context.Context, trackUUID uuid.UUID) (*domain.TrackPath, error) {
	trackPath := &domain.TrackPath{}

	err := c.conn.QueryRowContext(ctx, `
		SELECT (path, duration_ms, explicit)
		FROM tracks
		WHERE track_id = $1
		`,
		trackUUID,
	).Scan(&trackPath.Path, &trackPath.DurationMS, &trackPath.Explicit)
	switch err {
	case sql.ErrNoRows:
		return nil, domain.ErrTrackNotFound
	case nil:
		return trackPath, nil
	default:
		return nil, err
	}
}

func (c *PostgresCatalog) SearchTrack(ctx context.Context, trackName string, limit, offset int) ([]domain.Track, error) {
	rows, err := c.conn.QueryContext(ctx, `
		SELECT
			t.track_id, t.track_name, t.artist_id, ar.artist_name,
		 	t.album_id, al.album_name, t.explicit, t.created_at, t.duration_ms
		FROM tracks t
		JOIN artists ar ON ar.artist_id = t.artist_id
		JOIN albums al ON al.album_id = t.album_id
		WHERE t.track_name ILIKE '%' || $1 || '%' ESCAPE '\'
		ORDER BY t.listened DESC NULLS LAST
		LIMIT $2 OFFSET $3
		`,
		trackName, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tracks := []domain.Track{}
	for rows.Next() {
		var t domain.Track
		if err := rows.Scan(
			&t.TrackID, &t.TrackName, &t.ArtistID, &t.ArtistName,
			&t.AlbumID, &t.AlbumName, &t.Explicit, &t.CreatedAt, &t.DurationMS,
		); err != nil {
			return nil, err
		}
		tracks = append(tracks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tracks, nil

}

func (c *PostgresCatalog) TrackExists(ctx context.Context, trackUUID uuid.UUID) (bool, error) {
	var id uuid.UUID
	err := c.conn.QueryRowContext(ctx, `
		SELECT track_id
		FROM tracks
		WHERE track_id = $1
		`,
		trackUUID,
	).Scan(&id)
	switch err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

//////////////////////////////////////////////////////

func (c *PostgresCatalog) AddTrack(ctx context.Context, track *domain.Track) error {
	_, err := c.conn.ExecContext(ctx,
		`INSERT INTO tracks
			(track_id, track_name, artist_id, album_id,
			explicit, created_at, path, duration_ms)
		 VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		track.TrackID, track.TrackName, track.ArtistID, track.AlbumID, track.Explicit, track.CreatedAt, track.Path, track.DurationMS,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.ErrTrackAlreadyExists
		}
		return err
	}

	return nil
}

func (c *PostgresCatalog) IncrementListened(ctx context.Context, trackUUID uuid.UUID) error {
	_, err := c.conn.QueryContext(ctx, `
		UPDADE tracks
		SET listened = COALESCE(listened, 0) + 1
		WHERE track_id = $1
		`)

	switch err {
	case sql.ErrNoRows:
		return domain.ErrTrackNotFound
	case nil:
		return nil
	default:
		return err
	}
}
