module github.com/teper-ya-pomenyal/privy_stream/streaming_service

go 1.26.5

replace github.com/teper-ya-pomenyal/privy_stream/proto => ../proto

require (
	github.com/go-chi/chi/v5 v5.3.2
	github.com/google/uuid v1.6.0
	github.com/teper-ya-pomenyal/privy_stream/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.84.0
)

require (
	golang.org/x/net v0.57.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260706201446-f0a921348800 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
