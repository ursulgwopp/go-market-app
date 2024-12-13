# Use the official Golang image as the base image
FROM golang:1.22.4

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o ./bin/app ./cmd/main/main.go

# Expose port 1334 to the outside world
EXPOSE 1334

# Command to run the executable
CMD ["./bin/app"]