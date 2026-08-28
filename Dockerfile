FROM golang:1.26.5 AS builder
WORKDIR /app
COPY catalog_service/go.mod catalog_service/go.sum ./
COPY proto/ ./proto/
COPY catalog_service/ ./
RUN go build -o /app/bin/server ./cmd/server

FROM alpine:3.20
COPY --from=builder /app/bin/server /server
CMD ["/server"]