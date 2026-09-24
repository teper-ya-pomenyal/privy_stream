FROM golang:1.26.5 AS builder
WORKDIR /app
COPY user_service/go.mod user_service/go.sum ./user_service/
COPY proto/ ./proto/
COPY jwtmanager/ ./jwtmanager/
COPY user_service/ ./user_service/
WORKDIR /app/user_service
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server

FROM alpine:3.20
COPY --from=builder /app/server /server
CMD ["/server"]