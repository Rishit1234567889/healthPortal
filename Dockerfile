# Build stage
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o hospital-portal ./cmd/server

# Final stage
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/hospital-portal .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/migrations ./migrations


# Copy frontend static files
COPY --from=builder /app/public ./public



# Expose the application port
EXPOSE 8000

# Run the application
CMD ["./hospital-portal"]
