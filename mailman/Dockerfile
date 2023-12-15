# Use the official Golang image as the base image
FROM golang:1.21

WORKDIR /app
# Copy the Go application source code into the container
COPY mailman/ .
# Build the Go application
RUN go build -o main .

# Command to run the application
CMD ["./main"]