package main

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/trentenwen/grpc/cmd/server"
	pb "github.com/trentenwen/grpc/proto/addressbook"
	protos "github.com/trentenwen/grpc/proto/currency"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {

	/* Serialize data & Unserialize */

	book := &pb.AddressBook{
		People: []*pb.Person{
			{Name: "Alice", Age: 18},
			{Name: "Bob", Age: 26},
		},
	}

	data, err := proto.Marshal(book)

	if err != nil {
		fmt.Println("Fail to marshal address book", err)
	} else {
		fmt.Println(data)
	}

	var decodedBook pb.AddressBook
	err = proto.Unmarshal(data, &decodedBook)
	if err != nil {
		fmt.Println("Fail to unmarshal address book", err)
	} else {
		for _, person := range decodedBook.People {
			fmt.Println(person.Name, person.Age)
		}
	}

	/* Load gRPC service */

	log := hclog.Default()

	grpcServer := grpc.NewServer()
	currencyServer := server.NewCurrency(log)

	protos.RegisterCurrencyServer(grpcServer, currencyServer)
}
