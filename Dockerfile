# Use the official Golang image as a base
FROM golang:1.22.5

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules manifests and download dependencies
COPY go.mod ./

# Copy the entire project into the container
COPY . ./

# Build the Go application
RUN go build -o main ./cmd/server/main.go

# Expose the port your app runs on (e.g., 8080)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
