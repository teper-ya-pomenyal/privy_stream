CREATE TABLE playlists (
    playlist_id UUID PRIMARY KEY,
    playlist_name TEXT NOT NULL,
    owner_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);
