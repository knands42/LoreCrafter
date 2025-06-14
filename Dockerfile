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

# Copy the binary from the builder stage
COPY --from=builder /app/lorecrafter /app/lorecrafter

# Make sure the binaries are executable
RUN chmod +x /app/lorecrafter

# Expose the application port
EXPOSE 8000

# Run the application
CMD ["/app/lorecrafter"]
