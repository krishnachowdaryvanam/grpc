# gRPC Assignment

This repository contains a gRPC service that provides two endpoints to fetch user details based on user ID and a list of user IDs.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.20 or higher installed on your system.

## Getting Started

To get a local copy of this project and run the gRPC server, follow these steps:

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/yourusername/grpc_assignment.git
    ```

2. Change into the project directory:

    ```bash
    cd grpc_assignment
    ```

3. Build and run the gRPC server:

    ```bash
    docker build -t user-grpc-server .
    docker run -p 50051:50051 user-grpc-server
    ```

4. The gRPC server is now running on port 50051.

## Running the gRPC Client

To test the gRPC service using the client, follow these steps:

1. Open a new terminal window.

2. Change into the `client` directory:

    ```bash
    cd client
    ```

3. Run the client:

    ```bash
    go run client.go
    ```

The client will send requests to the gRPC server, and you should see the results on the terminal.

## Running Tests

To run tests for the gRPC service, you can use the following command:

```bash
go test ./...
