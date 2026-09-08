CREATE TABLE albums (
    album_id UUID PRIMARY KEY,
    artist_id UUID NOT NULL REFERENCES artists(artist_id),
    album_name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);
