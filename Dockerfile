# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/lorecrafter .

# Final stage
FROM alpine:latest

WORKDIR /app

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Create certificate cache directory
RUN mkdir -p /var/www/.cache
RUN chmod 755 /var/www/.cache

# Copy the binary from the builder stage
COPY --from=builder /app/lorecrafter /app/lorecrafter
COPY --from=builder /app/internal/adapter/database/migrations /app/internal/adapter/database/migrations

# Make sure the binary is executable
RUN chmod +x /app/lorecrafter
RUN mkdir /app/certs

# Expose the application port
EXPOSE 8000

# Run the application
CMD ["/app/lorecrafter", "--tls"]
