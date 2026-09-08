FROM golang:1.26.5 AS builder
WORKDIR /app
COPY gateway/go.mod gateway/go.sum ./gateway/
COPY proto/ ./proto/
COPY jwtmanager/ ./jwtmanager/
COPY gateway/ ./gateway/
WORKDIR /app/gateway
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/main.go

FROM alpine:3.20
COPY --from=builder /app/server /server
COPY --from=builder /app/gateway/keys /keys
CMD ["/server"]
