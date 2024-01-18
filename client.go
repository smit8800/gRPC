package main

import (
	"context"
	pb "gprc_practice/proto"
	"log"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	stime := time.Now()

	l := log.New(os.Stdout, "grpc-Client ", log.LstdFlags)

	conn, err := grpc.Dial("localhost:8010", grpc.WithInsecure())
	if err != nil {
		l.Fatalf("could not connect: %v", err)
	}
	client := pb.NewCalculatorClient(conn)

	var wg sync.WaitGroup

	// Example: Call the Add method
	wg.Add(1)
	go func() {
		defer wg.Done()
		addRequest := &pb.AddRequest{
			Num1: 10,
			Num2: 20,
		}
		res1, err := client.Add(context.Background(), addRequest)
		if err != nil {
			l.Printf("--> Add, RPC failed: %v", err)
		}
		l.Printf("--> Add, Response: %v", res1.Result)
	}()

	// Call StreamAdd in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		nums := []int32{10, 20, 30}
		stream1, err := client.StreamAdd(context.Background())
		if err != nil {
			l.Fatalf("--> StreamAdd, Failed to create stream: %v", err)
		}
		for _, num := range nums {
			req := &pb.StreamNumList{
				Nums: int32(num),
			}
			err := stream1.Send(req)
			l.Printf("--> StreamAdd, Send %d", num)
			if err != nil {
				l.Printf("--> StreamAdd, Failed to send request: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
		res2, err := stream1.CloseAndRecv()
		if err != nil {
			l.Printf("--> StreamAdd, Failed to receive response: %v", err)
		}
		l.Printf("--> StreamAdd, Received: %v", res2.Result)
	}()

	// Call AddStream in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		req := &pb.NumList{Nums: []int32{10, 20, 30}}
		stream2, err := client.AddStream(context.Background(), req)
		if err != nil {
			l.Printf("--> AddStream, Error calling AddStream: %v", err)
		}

		// Receive and print responses from the server stream.
		for {
			res3, err := stream2.Recv()
			if err != nil {
				l.Printf("--> AddStream, Error receiving response: %v", err)
				break
			}
			l.Printf("--> AddStream, Received: %d\n", res3.GetResult())
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	l.Printf("Client's been executed in %v", time.Since(stime))
}
