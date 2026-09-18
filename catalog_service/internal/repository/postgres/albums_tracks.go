package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

func (c *PostgresCatalog) AddTracksToAlbum(ctx context.Context, tracks []domain.AlbumTrack) error {
	if len(tracks) == 0 {
		return nil
	}

	albumID := tracks[0].AlbumUUID
	trackIDs := make([]uuid.UUID, 0, len(tracks))
	positions := make([]int, 0, len(tracks))
	for _, t := range tracks {
		if t.AlbumUUID != albumID {
			return domain.ErrInvalidUUID
		}
		trackIDs = append(trackIDs, t.TrackUUID)
		positions = append(positions, t.Position)
	}

	_, err := c.conn.ExecContext(ctx, `
		INSERT INTO albums_tracks (
			album_id,
			track_id,
			position
		)
		SELECT
			$1::uuid,
			input.track_id,
			input.position::integer
		FROM unnest(
			$2::uuid[],
			$3::integer[]
		) AS input(track_id, position)
	`,
		albumID, trackIDs, positions,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505": // unique_violation
				switch pgErr.ConstraintName {
				case "albums_tracks_pkey":
					return domain.ErrAlbumTrackAlreadyExists
				case "albums_tracks_album_id_position_key":
					return domain.ErrAlbumTrackPositionTaken
				default:
					return domain.ErrAlbumTrackAlreadyExists
				}
			case "23503": // foreign_key_violation
				switch pgErr.ConstraintName {
				case "albums_tracks_album_id_fkey":
					return domain.ErrAlbumNotFound
				case "albums_tracks_track_id_fkey":
					return domain.ErrTrackNotFound
				default:
					return err
				}
			case "23502": // not_null_violation
				return domain.ErrAlbumTrackInvalidPosition
			case "23514": // check_violation
				return domain.ErrAlbumTrackInvalidPosition
			}
		}
		return err
	}

	return nil
}
