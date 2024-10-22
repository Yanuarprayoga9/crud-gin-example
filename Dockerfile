# Use official Golang image
FROM golang:1.22.3-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app, specifying the main.go file path
RUN go build -o api ./cmd/.

# Expose the port
EXPOSE 8000

# Run the executable
CMD ["./api"]
