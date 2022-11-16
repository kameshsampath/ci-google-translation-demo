package server

import (
	"context"
	"io"
	"log"
	"testing"
	"time"

	"github.com/kameshsampath/ci-google-translation-demo/pkg/greeter"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var count = 0

func TestGreet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "grpc-service:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	req := &greeter.TranslationRequest{
		Message:     "Hello World",
		SourceLang:  "en",
		TargetLangs: []string{"ta"},
	}
	client := greeter.NewGreeterClient(conn)
	stream, err := client.Greet(context.TODO(), req)

	for {
		if err == io.EOF || count > 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		m, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		want := "வணக்கம் உலகம்"
		got := m.Message
		assert.Equalf(t, want, got, "Expecting %s but for %s", want, got)
		count++
	}
}
