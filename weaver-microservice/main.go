package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

const PORT = "12345"

var msg = createJSONPayload(500)
var requestBody = fmt.Sprintf("{\"id\": 123, \"action\": \"request_image\", \"message\": \"%s\"}", msg)
var wo = WorkObject{id: 123, action: "REQUEST_IMAGE", message: requestBody}

// app is the main component of the application. weaver.Run creates
// it and passes it to serve.
type app struct {
	weaver.Implements[weaver.Main]
	workerholic weaver.Ref[Workaholic]
}

func main() {
	fmt.Printf("Starting main service. Listening on localhost:%s...\n", PORT)

	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

//go:generate weaver generate ./...

func serve(ctx context.Context, app *app) error {
	// Get a network listener on address "localhost:12345".
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := app.Listener("hello", opts)
	if err != nil {
		return err
	}
	fmt.Printf("hello listener available on %v\n", lis)
	worker := app.workerholic.Get()

	// Serve the /hello endpoint.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, err := worker.Work(ctx, wo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, "Hello, World!")
	})

	return http.Serve(lis, nil)
}

func createJSONPayload(size int) (payload []byte) {
	letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for i := 0; i < size; i = i + 1 {
		letter := letters[rand.Intn(26)]
		payload = append(payload, byte(letter))
	}
	return
}
