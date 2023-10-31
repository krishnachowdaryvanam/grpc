# Specify the desired Go version (1.20)
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o userserver

# Expose the port on which the gRPC server will run
EXPOSE 50051

# Run the compiled binary
CMD ["./userserver"]
