package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
    if err := weaver.Run(context.Background(), serve); err != nil {
        log.Fatal(err)
    }
}

// app is the main component of the application. weaver.Run creates
// it and passes it to serve.
type app struct{
    weaver.Implements[weaver.Main]
    reverser weaver.Ref[Reverser]
}

// serve is called by weaver.Run and contains the body of the application.
func serve(ctx context.Context, app *app) error {
    opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
    lis, err := app.Listener("hello", opts)
    if err != nil {
        return err
    }
    fmt.Printf("hello listener available on %v\n", lis)

    // Serve the /hello endpoint.
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        if name == "" {
            name = "World"
        }
        reversed, err := app.reverser.Get().Reverse(ctx, name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Hello, %s!\n", reversed)
    })
    return http.Serve(lis, nil)
}
