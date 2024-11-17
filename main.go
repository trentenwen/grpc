package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/trentenwen/grpc/cmd/server"
	pb "github.com/trentenwen/grpc/protos/addressbook"
	protos "github.com/trentenwen/grpc/protos/currency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	reflection.Register(grpcServer)

	protos.RegisterCurrencyServer(grpcServer, currencyServer)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	grpcServer.Serve(l)
}
