# Stage 1: Build the Go application
FROM golang:1.21.4 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the host machine to the container
COPY . .

# Ensure dependencies are up to date
RUN go mod tidy

# Compile the binary for Linux explicitly (important for Docker)
RUN GOOS=linux GOARCH=amd64 go build -o receipt-processor .

# Stage 2: Create a lightweight final image
FROM alpine:latest

# Set the working directory inside the final container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/receipt-processor .

# Ensure the binary is executable
RUN chmod +x receipt-processors

# Expose port 8080 for API access
EXPOSE 8080

# Command to run the application
CMD ["./receipt-processor"]
