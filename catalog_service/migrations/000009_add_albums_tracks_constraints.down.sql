ALTER TABLE albums_tracks
    DROP CONSTRAINT albums_tracks_album_id_position_key,
    DROP CONSTRAINT albums_tracks_position_check,
    ALTER COLUMN position DROP NOT NULL;
