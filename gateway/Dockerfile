FROM golang:1.26.5 AS builder
WORKDIR /app
COPY gateway/go.mod gateway/go.sum ./
COPY proto/ ./proto
COPY gateway/ ./
RUN go build -o /app/bin/server ./cmd/server

FROM alpine:3.20
COPY --from=builder /app/bin/server /server
CMD ["/server"]