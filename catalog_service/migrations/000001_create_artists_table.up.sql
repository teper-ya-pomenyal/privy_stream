CREATE TABLE artists (
    artist_id UUID PRIMARY KEY,
    artist_name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
