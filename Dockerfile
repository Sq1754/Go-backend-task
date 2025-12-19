# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (needed for go mod)
RUN apk add --no-cache git

# Copy go mod files first (cache optimization)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o server ./cmd/server

# ---------- Runtime Stage ----------
FROM alpine:latest

WORKDIR /app

# Copy compiled binary
COPY --from=builder /app/server .

# Expose Fiber port
EXPOSE 3000

# Run server
CMD ["./server"]
