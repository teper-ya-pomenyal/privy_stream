FROM golang:1.26.5 AS builder
WORKDIR /app
COPY user_service/go.mod user_service/go.sum ./
COPY proto/ ./proto
COPY user_service/ ./
RUN go build -o /app/bin/server ./cmd/server

FROM alpine:3.20
COPY --from=builder /app/bin/server /server
CMD ["/server"]