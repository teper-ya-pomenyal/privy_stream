module github.com/teper-ya-pomenyal/privy_stream/user_service

go 1.26.5

require (
	github.com/google/uuid v1.6.0
	github.com/jackc/pgx/v5 v5.10.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/redis/go-redis/v9 v9.22.0
	github.com/teper-ya-pomenyal/privy_stream/jwtmanager v0.0.0-20260830100439-1ca989f99171
	github.com/teper-ya-pomenyal/privy_stream/proto v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.56.0
	google.golang.org/grpc v1.83.2
	google.golang.org/protobuf v1.36.11
)

replace github.com/teper-ya-pomenyal/privy_stream/jwtmanager => ../jwtmanager

replace github.com/teper-ya-pomenyal/privy_stream/proto => ../proto

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/lib/pq v1.12.3 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa // indirect
)
