FROM golang:1.26.5 AS builder
WORKDIR /app
COPY streaming_service/go.mod streaming_service/go.sum ./streaming_service/
COPY proto/ ./proto/
COPY streaming_service/ ./streaming_service/
WORKDIR /app/streaming_service
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server ./cmd/main.go

FROM alpine:3.20
COPY --from=builder /app/bin/server /server
CMD ["/server"]