package server

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"

	"github.com/kameshsampath/ci-google-translation-demo/pkg/greeter"
)

var (
	once   sync.Once
	client *translate.Client
	err    error
)

// LinguaGreeterServer implements the greeter.Greeter service
type LinguaGreeterServer struct {
	greeter.UnimplementedGreeterServer
}

// Greet implements greeter.GreeterServer
func (*LinguaGreeterServer) Greet(req *greeter.TranslationRequest, stream greeter.Greeter_GreetServer) error {
	// keeps streaming the translations by rotating the languages
	for {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(req.TargetLangs), func(i, j int) {
			req.TargetLangs[i], req.TargetLangs[j] = req.TargetLangs[j], req.TargetLangs[i]
		})
		tl := req.TargetLangs[0]
		t, err := client.Translate(context.Background(),
			[]string{req.Message},
			language.MustParse(tl),
			&translate.Options{
				Source: language.MustParse(req.SourceLang),
				Format: translate.Text,
			})
		if err != nil {
			return err
		}
		if len(t) > 0 {
			if err := stream.Send(&greeter.TranslationReply{
				Message: t[0].Text,
				Lang:    tl,
			}); err != nil {
				return err
			}
		}
		time.Sleep(3 * time.Second)
	}
}

var _ greeter.GreeterServer = (*LinguaGreeterServer)(nil)

func init() {
	if apiKey, ok := os.LookupEnv("DEMO_API_KEY"); ok {
		ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
		defer cancel()
		client, err = translate.NewClient(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			panic(err)
		}
	} else {
		panic(fmt.Errorf("no api key found"))
	}
}
