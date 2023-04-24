package main

import (
	"context"
	"log"
	"time"

	"github.com/ivan-sabo/protobuf/examples/go/apiv1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	client()
}

func client() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(
		"localhost:50051",
		opts...,
	)
	if err != nil {
		log.Fatalf("an error occured: %e", err)
	}

	client := apiv1.NewAddressBookServiceClient(conn)
	req := apiv1.AddressBookRequest{Id: 2}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetAddressBook(ctx, &req)
	if err != nil {
		log.Fatalf("an error occured: %e", err)
	}

	log.Printf("response: %v", resp)
}
