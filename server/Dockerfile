FROM golang:1.24.2-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o server main.go
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata
RUN adduser -D -u 1000 -g '' appuser
COPY --from=builder --chown=appuser:appuser /app/server /app/server
USER appuser
EXPOSE 8080
CMD ["/app/server"]