# Build stage
FROM golang:1.25-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/app/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates in case the app needs to make HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
# Copy config files if needed (assuming they are in configs/ or at root)
COPY --from=builder /app/configs ./configs
# Copy .env.example as template if needed, or if user wants .env in container (usually provided via env vars)
# COPY --from=builder /app/.env.example ./.env

# Command to run the executable
ENTRYPOINT ["./main"]
