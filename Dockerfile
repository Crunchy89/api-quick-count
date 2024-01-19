# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy only the go.mod and go.sum files to leverage Docker cache
COPY go.mod go.sum ./

# Download and verify dependencies
RUN go mod download && go mod verify

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -v -o /usr/local/bin/api

# Expose port 80
EXPOSE 80/tcp

# Run the executable
CMD ["api"]
