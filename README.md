"""mermaid
flowchart TB
    Client[Web-клиент] -->|HTTP/JSON| Gateway[API Gateway]
    Gateway -->|gRPC| User[User Service]
    Gateway -->|gRPC| Catalog[Catalog Service]
    Gateway -->|HTTP reverse proxy| Streaming[Streaming Service]
    Streaming -->|gRPC| Catalog

    User --> PGUsers[(Postgres: users_db)]
    Catalog --> PGCatalog[(Postgres: catalog_db)]
    User --> RedisSessions[(Redis: users_sessions)]
    Catalog --> RedisCache[(Redis: catalog_cache)]
    Streaming --> Storage[(File Storage)]
"""

"""mermaid
erDiagram
    USERS {
        uuid id_user PK
        string user_name
        string password_hash
        date birth_date
        timestamp created_at
    }

    ARTISTS {
        uuid id_artist PK
        string artist_name
        timestamp created_at
    }

    ALBUMS {
        uuid id_album PK
        uuid id_artist FK
        string album_name
        timestamp created_at
    }

    TRACKS {
        uuid id_track PK
        uuid artist_id FK
        uuid album_id FK
        string track_name
        bool explicit
        string path
        int duration_ms
        timestamp created_at
    }

    PLAYLISTS {
        uuid id_playlist PK
        string playlist_name
        uuid owner_id
        timestamp created_at
    }

    PLAYLIST_TRACKS {
        uuid id_playlist PK,FK
        uuid id_track PK,FK
        int position
        timestamp added_at
    }

    ARTISTS ||--o{ ALBUMS : "has"
    ARTISTS ||--o{ TRACKS : "performs"
    ALBUMS ||--o{ TRACKS : "contains"
    PLAYLISTS ||--o{ PLAYLIST_TRACKS : "includes"
    TRACKS ||--o{ PLAYLIST_TRACKS : "appears in"
"""