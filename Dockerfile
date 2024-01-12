# Start from the latest golang base image as the build stage
FROM golang:1.18 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY *.go ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o thttp .

# Start a new stage from scratch
FROM alpine:latest

# Copy the executable from the builder stage
COPY --from=builder /app/thttp /usr/local/bin/thttp

# Expose port 8000 to the outside
EXPOSE 8000

# Command to run the executable
CMD ["thttp"]