package main

import (
	"context"
	"github.com/ShatteredRealms/Characters/pkg/characterspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:8081", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := characterspb.NewCharacterServiceClient(conn)

	log.Printf("Sending message...")

	_, err = client.Create(context.Background(), &characterspb.CreateCharacterMessage{})

	log.Printf("Message received: %v", err)
}
