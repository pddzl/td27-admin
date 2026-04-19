# syntax=docker/dockerfile:1
FROM golang:1.25-alpine AS builder

WORKDIR /server

# Install git for private modules (if needed) and ca-certificates
RUN apk add --no-cache git ca-certificates tzdata

# Download deps first (better layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o td27Server ./cmd/server/

# Final stage
FROM alpine:3.19

LABEL maintainer="pddzl5@foxmail.com"

# Security: non-root user
RUN addgroup -g 1000 -S td27 && adduser -u 1000 -S td27 -G td27

WORKDIR /server

# Install ca-certificates for HTTPS calls
RUN apk add --no-cache ca-certificates

# Create necessary directories
RUN mkdir -p ./configs ./resource/upload && chown -R td27:td27 .

# Copy artifacts
COPY --from=builder --chown=td27:td27 /server/td27Server ./
COPY --from=builder --chown=td27:td27 /server/configs/config.yaml ./configs/
COPY --from=builder --chown=td27:td27 /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai

USER td27

EXPOSE 8888

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:8888/health || exit 1

CMD ["./td27Server"]
