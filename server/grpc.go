package main

import (
	"context"
	"log"
	"net"

	"github.com/ivan-sabo/protobuf/examples/go/apiv1"
	"google.golang.org/grpc"
)

func main() {
	TestGRPC()
}

func TestGRPC() {
	s := grpc.NewServer()
	testServer := &TestGRPCServer{}

	apiv1.RegisterAddressBookServiceServer(s, testServer)

	l, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("an error occured: %e", err)
	}

	log.Printf("server listening at %v", l.Addr())

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("an error occured: %e", err)
	}
}

type TestGRPCServer struct {
	apiv1.UnimplementedAddressBookServiceServer
}

func (s *TestGRPCServer) GetAddressBook(ctx context.Context, r *apiv1.AddressBookRequest) (*apiv1.AddressBookResponse, error) {
	log.Printf("Request arrived, address book id: %d", r.Id)

	return &apiv1.AddressBookResponse{
		Addressbook: &apiv1.AddressBook{
			People: []*apiv1.Person{
				{
					Id:    666,
					Email: "mark@google.com",
					Phones: []*apiv1.Person_PhoneNumber{
						{
							Number: "+39333111",
							Type:   apiv1.Person_WORK,
						},
					},
				},
			},
		},
	}, nil
}
