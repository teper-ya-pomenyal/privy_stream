FROM golang:1.26.5 AS builder
WORKDIR /app
COPY gateway/go.mod ./
COPY proto/ ./proto
COPY gateway/ ./
RUN go build -o /app/bin/server ./cmd/main.go

FROM alpine:3.20
COPY --from=builder /app/bin/server /server
CMD ["/server"]