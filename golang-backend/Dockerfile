# Use an official Go image as the base image
FROM golang:1.22 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to take advantage of Docker caching
COPY go.mod go.sum ./

# Download dependencies (this will be cached if go.mod/go.sum haven't changed)
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Use a lightweight image for the final container
FROM gcr.io/distroless/base

# Set the working directory inside the container
WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
