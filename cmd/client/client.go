package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/trentenwen/grpc/protos/currency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Client setup...")

	serverAddr := "localhost:9092"
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.NewClient(serverAddr, opts)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(0)
	}

	client := pb.NewCurrencyClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.RateRequest{
		Base:        pb.Currencies_EUR,
		Destination: pb.Currencies_USD,
	}

	resp, err := client.GetRate(ctx, req)
	if err != nil {
		fmt.Println("GetRate Error", err)
		os.Exit(0)
	}

	fmt.Println("Success", resp.Rate)
}
