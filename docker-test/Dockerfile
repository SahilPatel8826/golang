# Step 1: Build the Go app
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./

RUN go mod download

# Copy the whole project
COPY . .

# Build the app
RUN go build -o myapp .

# Step 2: Run the app with a lightweight image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from builder stage
COPY --from=builder /app/myapp .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./myapp"]
