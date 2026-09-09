CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE INDEX idx_tracks_track_name_trgm
ON tracks USING GIN (track_name gin_trgm_ops);
