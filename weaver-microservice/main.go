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

// app is the main component of the application. weaver.Run creates
// it and passes it to serve.
type app struct {
	weaver.Implements[weaver.Main]
	workerholic weaver.Ref[Worker]
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

	http.HandleFunc("/thick", func(w http.ResponseWriter, r *http.Request) {
		_, err := worker.Work(ctx, wo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, "Hello, World!")
	})

	http.HandleFunc("/long", func(w http.ResponseWriter, r *http.Request) {
		_, err := worker.Work2(ctx, wo2)
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

var msg = createJSONPayload(500)
var wo = WorkObject{id: 123, action: "REQUEST_IMAGE", message: string(msg)}
var wo2 = WorkObject2{id: 123, action: "REQUEST_USER",
	message1:  string(createJSONPayload(10)),
	message2:  string(createJSONPayload(10)),
	message3:  string(createJSONPayload(10)),
	message4:  string(createJSONPayload(10)),
	message5:  string(createJSONPayload(10)),
	message6:  string(createJSONPayload(10)),
	message7:  string(createJSONPayload(10)),
	message8:  string(createJSONPayload(10)),
	message9:  string(createJSONPayload(10)),
	message10: string(createJSONPayload(10)),
	message11: string(createJSONPayload(10)),
	message12: string(createJSONPayload(10)),
	message13: string(createJSONPayload(10)),
	message14: string(createJSONPayload(10)),
	message15: string(createJSONPayload(10)),
	message16: string(createJSONPayload(10)),
	message17: string(createJSONPayload(10)),
	message18: string(createJSONPayload(10)),
	message19: string(createJSONPayload(10)),
	message20: string(createJSONPayload(10)),
}
