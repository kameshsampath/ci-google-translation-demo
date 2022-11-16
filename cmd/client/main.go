package main

import (
	"context"
	"io"
	"log"

	"github.com/kameshsampath/ci-google-translation-demo/pkg/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	req := &greeter.TranslationRequest{
		Message:     "Hello World",
		SourceLang:  "en",
		TargetLangs: []string{"ta", "kn", "te"},
	}
	client := greeter.NewGreeterClient(conn)
	stream, err := client.Greet(context.TODO(), req)

	for {
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		m, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf(`Translated "%s" from "%s" to "%s" as "%s"`, req.Message, req.SourceLang, m.Lang, m.Message)
	}
}
