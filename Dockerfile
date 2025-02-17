# Build stage
FROM golang:1.22.3-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o oil_selection main.go

# Final stage
FROM ubuntu:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/oil_selection .



# Expose port 8000
EXPOSE 8000

# Run the binary
CMD ["./oil_selection"]
