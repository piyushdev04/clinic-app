FROM golang:1.22-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app
RUN go build -o server ./cmd/server

# Expose port 8080
EXPOSE 8080

# Start the server
CMD ["./server"] 