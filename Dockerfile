# Stage 1: Build
FROM golang:1.23-alpine AS builder

# Install CA certificates and UPX for compression
RUN apk add --no-cache ca-certificates upx

WORKDIR /app

# Module Cache Optimization
COPY go.mod go.sum* ./
RUN go mod download

COPY . .

# Build with stripping flags for smaller binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
    -ldflags="-s -w -extldflags '-static'" \
    -o mcp-server main.go

# Extreme compression with UPX
RUN upx -9 mcp-server

# Stage 2: Final (Zero-size base)
FROM scratch

# Copy certificates for external API calls
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

# Copy only the compiled binary
COPY --from=builder /app/mcp-server .

# Run the binary
ENTRYPOINT ["./mcp-server"]
