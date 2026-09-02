FROM golang:1.26.5 AS builder
WORKDIR /app
COPY go.work go.work.sum ./
COPY user_service/go.mod user_service/go.sum ./user_service/
COPY proto/ ./proto/
COPY jwtmanager/ ./jwtmanager/
COPY user_service/ ./user_service/
WORKDIR /app/user_service
RUN go build -o /app/bin/server ./cmd/server

FROM alpine:3.20
COPY --from=builder /app/bin/server /server
CMD ["/server"]