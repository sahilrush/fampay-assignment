# # Stage 1: Build the Go binary
# FROM golang:1.23-alpine AS builder

# # Install git for fetching dependencies
# RUN apk add --no-cache git

# # Set working directory
# WORKDIR /app

# # Copy go mod files
# COPY go.mod go.sum ./

# # Download dependencies
# RUN go mod download

# # Copy source code
# COPY . .

# # Build the application for the correct architecture (amd64)
# RUN GOARCH=amd64 GOOS=linux go build -o main -ldflags="-s -w" .

# # Stage 2: Create the final image
# FROM alpine:3.19

# # Install necessary runtime dependencies
# RUN apk add --no-cache libc6-compat

# # Set working directory
# WORKDIR /app

# # Copy the binary from the builder stage
# COPY --from=builder /app/main .

# # Expose port 8080
# EXPOSE 8080

# # Run the application
# CMD ["./main"]



FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "main.go"]