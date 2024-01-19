package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	pb "gprc_practice/proto"
	proto "gprc_practice/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// calculatorServer implements the Calculator service defined in the proto package.
type calculatorServer struct {
	proto.UnimplementedCalculatorServer
}

// NewCalculatorServer creates a new instance of the calculatorServer.
func NewCalculatorServer() *calculatorServer {
	return &calculatorServer{}
}

// Add implements the Add RPC method of the Calculator service.
func (s *calculatorServer) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	l := log.New(os.Stdout, "grpc-Server --> Add ", log.LstdFlags)
	result := req.Num1 + req.Num2

	l.Printf("Sending response: %d", result)
	return &proto.AddResponse{Result: result}, nil
}

func (s *calculatorServer) StreamAdd(stream pb.Calculator_StreamAddServer) error {
	l := log.New(os.Stdout, "grpc-Server --> StreamAdd ", log.LstdFlags)
	_sum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			l.Printf("Send: %d", _sum)
			return stream.SendAndClose(&proto.AddResponse{Result: _sum})
		}
		if err != nil {
			return err
		}
		l.Printf("Recieved: %d", req.Nums)
		_sum += int32(req.Nums)
	}

}

func (s *calculatorServer) AddStream(req *proto.NumList, stream pb.Calculator_AddStreamServer) error {
	l := log.New(os.Stdout, "grpc-Server --> AddStream ", log.LstdFlags)
	_sum := 0
	for _, num := range req.GetNums() {
		_sum += int(num)
		res := &pb.StreamAddResponse{
			Result: int32(_sum),
		}
		err := stream.Send(res)
		if err != nil {
			l.Fatalf("Error in Sending")
			return err
		}
		l.Printf("Sending: %d", _sum)
		// 2-second delay to simulate a long-running process
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *calculatorServer) Bi_Add(stream pb.Calculator_Bi_AddServer) error {
	l := log.New(os.Stdout, "grpc-Server --> Bi ", log.LstdFlags)
	_sum := int32(0)

	for {
		// reciever: stream data
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		l.Printf("Recieved: %d", req.Nums)

		// sending: on stream
		_sum += int32(req.Nums)
		res := &pb.StreamAddResponse{
			Result: int32(_sum),
		}
		err = stream.Send(res)
		if err != nil {
			l.Fatalf("Error in Sending")
			return err
		}
		l.Printf("Sending: %d", _sum)
	}
}

func main() {
	// Create a logger for the server.
	l := log.New(os.Stdout, "grpc-Server ", log.LstdFlags)

	// Create a new gRPC server.
	grpc_server := grpc.NewServer()

	// Create an instance of the calculatorServer.
	calculator_server := NewCalculatorServer()

	// Register the Calculator service with the gRPC server.
	proto.RegisterCalculatorServer(grpc_server, calculator_server)

	// Enable reflection for grpcurl and other tools (for gRPCCurl)
	reflection.Register(grpc_server)

	lis, err := net.Listen("tcp", ":8010")
	if err != nil {
		l.Fatalf("failed to listen: %v", err)
	}
	l.Println("gRPC server started on :8010")
	if err := grpc_server.Serve(lis); err != nil {
		l.Fatalf("failed to serve: %v", err)
	}
	grpc_server.GracefulStop()

	// Start the gRPC server in a separate goroutine.
	go func() {
		lis, err := net.Listen("tcp", ":8010")
		if err != nil {
			l.Fatalf("failed to listen: %v", err)
		}
		l.Println("gRPC server started on :8010")
		if err := grpc_server.Serve(lis); err != nil {
			l.Fatalf("failed to serve: %v", err)
		}
	}()

	// Set up a signal channel to capture interrupt and kill signals.
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Wait for a signal to initiate graceful shutdown.
	sig := <-sigChan

	// Log the received signal and initiate graceful shutdown of the gRPC server.
	l.Println("received terminate, graceful shutdown", sig)
	grpc_server.GracefulStop()
}
