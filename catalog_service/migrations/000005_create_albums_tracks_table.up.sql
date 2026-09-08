CREATE TABLE albums_tracks (
    album_id UUID NOT NULL REFERENCES albums(album_id),
    track_id UUID NOT NULL REFERENCES tracks(track_id),
    position INTEGER,
    PRIMARY KEY (album_id, track_id)
);
