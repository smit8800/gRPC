# gRP Implementation

This repository contains a practice implementation of gRPC (Google Remote Procedure Call) in Go, showcasing various RPC methods including unary, server streaming, and client streaming.

## Features

### Unary RPC
- **Add Method:** Performs a unary RPC by adding two numbers and returning the result.

### Server Streaming RPC
- **StreamAdd Method:** Demonstrates server streaming, where the server receives a stream of numbers and returns the cumulative sum.

### Client Streaming RPC
- **AddStream Method:** Showcases client streaming, with the client sending a stream of numbers and the server returning the intermediate sums.

## Getting Started

Follow these instructions to run the gRPC server and client on your local machine.

### Prerequisites

- [Go](https://golang.org/) installed on your machine.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/grpc-practice.git
    ```

2. Navigate to the project directory:

    ```bash
    cd grpc-practice
    ```

### Running the Server

```bash
go run main.go
