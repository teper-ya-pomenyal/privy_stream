package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

func (c *PostgresCatalog) GetPlaylistByID(ctx context.Context, playlistUUID uuid.UUID) (*domain.Playlist, error) {
	var p domain.Playlist
	err := c.pool.QueryRow(ctx, `
		SELECT playlist_id, playlist_name, owner_id, created_at
		FROM playlists
		WHERE playlist_id = $1
		`,
		playlistUUID,
	).Scan(&p.PlaylistID, &p.PlaylistName, &p.OwnerID, &p.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrPlaylistNotFound
		}
		return nil, err
	}

	p.Tracks, err = c.GetPlaylistTracks(ctx, playlistUUID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *PostgresCatalog) GetPlaylistTracks(ctx context.Context, playlistUUID uuid.UUID) ([]domain.LightPlaylistTrack, error) {
	rows, err := c.pool.Query(ctx, `
		SELECT pt.track_id, t.track_name, t.explicit, t.duration_ms, pt.position
		FROM playlists_tracks pt
		JOIN tracks t ON t.track_id = pt.track_id
		WHERE pt.playlist_id = $1
		ORDER BY pt.position ASC
		`,
		playlistUUID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tracks := []domain.LightPlaylistTrack{}
	for rows.Next() {
		var t domain.LightPlaylistTrack
		if err := rows.Scan(&t.TrackID, &t.TrackName, &t.Explicit, &t.DurationMS, &t.Position); err != nil {
			return nil, err
		}
		tracks = append(tracks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tracks, nil
}
