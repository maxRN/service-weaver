package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	worker weaver.Ref[Worker]
}

func main() {
	fmt.Printf("Starting main service. Listening on localhost:%s...\n", PORT)

	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

// serve is called by weaver.Run and contains the body of the application.
func serve(ctx context.Context, app *app) error {
	opts := weaver.ListenerOptions{LocalAddress: "localhost:" + PORT}
	lis, err := app.Listener("hello", opts)
	if err != nil {
		return err
	}
	fmt.Printf("hello listener available on %v\n", lis)

	worker := app.worker.Get()
	// Serve the /hello endpoint.
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		// sizeString := r.URL.Query().Get("size")
		// size, _ := strconv.Atoi(sizeString)
		// timesString := r.URL.Query().Get("times")
		// times, _ := strconv.Atoi(timesString)

		// Get worker reference and do work
		doWork(1, 1, ctx, worker)

		fmt.Fprint(w, "Hello, World!")
	})
	return http.Serve(lis, nil)
}

func doWork(size int, times int, ctx context.Context, worker Worker) {
	_, err := worker.Work(ctx, wo)
	if err != nil {
		log.Println(err)
	}
}

func createJSONPayload(size int) (payload []byte) {
	letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for i := 0; i < size; i = i + 1 {
		letter := letters[rand.Intn(26)]
		payload = append(payload, byte(letter))
	}
	return
}
