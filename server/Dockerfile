# Stage 1: Build the Go binary
FROM golang:1.24.2-alpine AS builder
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git ca-certificates build-base

# Download dependencies first (for better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -extldflags '-static'" \
    -a -installsuffix cgo \
    -o server main.go

# Stage 2: Minimal runtime image
FROM alpine:latest
WORKDIR /app

# Install CA certificates for HTTPS requests
RUN apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates

# Create non-root user
RUN adduser -D -u 1000 -g '' appuser

# Copy the binary from builder
COPY --from=builder --chown=appuser:appuser /app/server /app/server
COPY --from=builder --chown=appuser:appuser /app/.env /app/.env

# Switch to non-root user
USER appuser

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Expose port
EXPOSE 8080

# Run the binary
CMD ["/app/server"]