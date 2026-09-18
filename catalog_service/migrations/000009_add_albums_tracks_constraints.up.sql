ALTER TABLE albums_tracks
    ALTER COLUMN position SET NOT NULL,
    ADD CONSTRAINT albums_tracks_position_check CHECK (position > 0),
    ADD CONSTRAINT albums_tracks_album_id_position_key UNIQUE (album_id, position);
