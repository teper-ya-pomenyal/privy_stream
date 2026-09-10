package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

func (c *PostgresCatalog) SearchArtist(ctx context.Context, artistName string, limit, offset int) ([]domain.LightArtist, error) {
	rows, err := c.conn.QueryContext(ctx, `
		SELECT artist_id, artist_name
		FROM artists
		WHERE artist_name ILIKE '%' || $1 || '%'
		ORDER BY artist_name ASC
		LIMIT $2 OFFSET $3
		`,
		artistName, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	artists := []domain.LightArtist{}
	for rows.Next() {
		var a domain.LightArtist
		if err := rows.Scan(
			&a.ArtistID, &a.ArtistName,
		); err != nil {
			return []domain.LightArtist{}, err
		}
		artists = append(artists, a)
	}
	if err = rows.Err(); err != nil {
		return []domain.LightArtist{}, err
	}
	return artists, nil
}

func (c *PostgresCatalog) GetArtistAlbums(ctx context.Context, artistUUID uuid.UUID, limit, offset int32) ([]domain.LightAlbum, error) {
	rows, err := c.conn.QueryContext(ctx, `
			SELECT album_id, album_name, created_at
			FROM albums
			WHERE artist_id = $1
			ORDER BY created_at DESC
			LIMIT $2 OFFSET $3
			`,
		artistUUID, limit, offset,
	)
	if err != nil {
		return []domain.LightAlbum{}, err
	}
	defer rows.Close()

	albums := []domain.LightAlbum{}

	for rows.Next() {
		var a domain.LightAlbum
		if err := rows.Scan(&a.AlbumUUID, &a.AlbumName, &a.CreatedAt); err != nil {
			return []domain.LightAlbum{}, err
		}
		albums = append(albums, a)
	}
	if err := rows.Err(); err != nil {
		return []domain.LightAlbum{}, err
	}
	return albums, nil
}

func (c *PostgresCatalog) GetArtistTracks(ctx context.Context, artistUUID uuid.UUID, limit, offset int32) ([]domain.LightTrack, error) {
	rows, err := c.conn.QueryContext(ctx, `
		SELECT track_id, track_name, explicit, duration_ms
		FROM tracks
		WHERE artist_id = $1
		ORDER BY listened DESC
		LIMIT $2 OFFSET $3
		`,
		artistUUID, limit, offset,
	)
	if err != nil {
		return []domain.LightTrack{}, err
	}
	defer rows.Close()
	tracks := []domain.LightTrack{}

	for rows.Next() {
		t := domain.LightTrack{}
		if err := rows.Scan(&t.TrackID, &t.TrackName, &t.Explicit, &t.DurationMS); err != nil {
			return []domain.LightTrack{}, err
		}
		tracks = append(tracks, t)
	}
	if err = rows.Err(); err != nil {
		return []domain.LightTrack{}, err
	}
	return tracks, nil
}
