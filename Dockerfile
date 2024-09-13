# Use the official Go image as a build stage
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY src/go.mod ./  
RUN go mod download

# Copy the source code
COPY src/ .
COPY . .         

# Build the Go application
RUN go build -o main .

# Use a minimal base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .
# Copy the templates directory from the build stage
COPY --from=builder /app/templates ./templates

# Expose the port the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]

