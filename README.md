gRPC Practice Repository
This repository contains a practice implementation of gRPC (Google Remote Procedure Call) in Go, showcasing various RPC methods including unary, server streaming, and client streaming.

Features
Unary RPC
Add Method: Performs a unary RPC by adding two numbers and returning the result.
Server Streaming RPC
StreamAdd Method: Demonstrates server streaming, where the server receives a stream of numbers and returns the cumulative sum.
Client Streaming RPC
AddStream Method: Showcases client streaming, with the client sending a stream of numbers and the server returning the intermediate sums.
Getting Started
Follow these instructions to run the gRPC server and client on your local machine.

Prerequisites
Go installed on your machine.
Installation
Clone the repository:

bash
Copy code
git clone https://github.com/yourusername/grpc-practice.git
Navigate to the project directory:

bash
Copy code
cd grpc-practice
Running the Server
bash
Copy code
go run main.go
Running the Client
Open a new terminal and run:

bash
Copy code
go run client.go
Protobuf Definitions
The gRPC services and message types are defined in the calculator.proto file. These include:

Add: Unary RPC to add two numbers.
StreamAdd: Server streaming RPC for cumulative sum.
AddStream: Client streaming RPC for intermediate sums.
License
This project is licensed under the MIT License - see the LICENSE file for details.

Acknowledgments
gRPC - Google's open-source RPC framework.
Protocol Buffers - Language-agnostic data serialization.
Feel free to explore and modify the code to understand and experiment with gRPC concepts.

