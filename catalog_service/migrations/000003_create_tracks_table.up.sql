CREATE TABLE tracks (
    track_id UUID PRIMARY KEY,
    track_name TEXT NOT NULL,
    artist_id UUID NOT NULL REFERENCES artists(artist_id),
    album_id UUID NOT NULL REFERENCES albums(album_id),
    explicit BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);
