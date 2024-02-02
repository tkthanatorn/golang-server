# Use a golang base image
FROM golang:1.21.3-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to the container
COPY go.mod go.sum ./

# Download dependencies (including private modules)
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Use a minimal Alpine-based image as the base for the final image
FROM alpine:latest

# Copy the built binary from the previous stage
COPY --from=builder /app/app /usr/local/bin/app

# Expose the port your application listens on
EXPOSE 8080

# Set the command to start your application
CMD ["/usr/local/bin/app"]