package main

import (
	"fmt"
	junk "github.com/xtreme-jesse-malone/grpc-streaming-example/proto"
	"google.golang.org/grpc"
	"io"
	"net"
)

type junkserver struct {
	junk.UnimplementedJunkServer
}

func main() {
	fmt.Printf("starting")
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Printf("error starting listener")
	}
	server := grpc.NewServer()
	junk.RegisterJunkServer(server, &junkserver{})

	if err := server.Serve(lis); err != nil {
		fmt.Printf("Failed serving %v", err)
	}
}

func (*junkserver) TakeJunk(stream junk.Junk_TakeJunkServer) error {
	for {
		thing, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&junk.EmptyMsg{})
		}
		stuff := thing.GetJunk()
		if stuff != "" {
			fmt.Printf("Got: %v\n", thing.GetJunk())
		}
	}
}
