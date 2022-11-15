package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/kameshsampath/ci-google-translation-demo/pkg/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func lingua_greeter(client greeter.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.Greet(ctx, &greeter.TranslationRequest{
		Message:     "Hello World",
		SourceLang:  "en",
		TargetLangs: []string{"ta", "kn", "te"},
	})

	if err != nil {
		log.Fatal(err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s", res.Message)
	}

}

func main() {
	conn, err := grpc.Dial(
		"greeter-service:9090",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := greeter.NewGreeterClient(conn)
	lingua_greeter(client)
}
