FROM golang:1.26.5 AS builder
WORKDIR /app
COPY catalog_service/go.mod catalog_service/go.sum ./catalog_service/
COPY proto/ ./proto/
COPY catalog_service/ ./catalog_service/
WORKDIR /app/catalog_service
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/main.go

FROM alpine:3.20
COPY --from=builder /app/server /server
CMD ["/server"]
