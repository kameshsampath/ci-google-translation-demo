package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/kameshsampath/ci-google-translation-demo/pkg/greeter"
	impl "github.com/kameshsampath/ci-google-translation-demo/pkg/server"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, &impl.LinguaGreeterServer{})
	log.Printf("Server listening on %s", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
