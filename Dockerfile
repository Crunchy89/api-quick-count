# Use the official Golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go build -o myapp .

# Expose the port that the application will run on
EXPOSE 5678

# Command to run the application
CMD ["./myapp"]
