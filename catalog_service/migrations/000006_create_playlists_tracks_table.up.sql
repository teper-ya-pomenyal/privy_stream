CREATE TABLE playlists_tracks (
    playlist_id UUID NOT NULL REFERENCES playlists(playlist_id),
    track_id UUID NOT NULL REFERENCES tracks(track_id),
    position INTEGER NOT NULL,
    added_at TIMESTAMPTZ,
    PRIMARY KEY (playlist_id, track_id)
);
