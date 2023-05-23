package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
    _ "net/http/pprof"
)

const PORT = "8069"

func main() {
	fmt.Printf("Starting work service. Listening on localhost:%s\n", PORT)

	// Start listener
	http.Handle("/do-work", http.HandlerFunc(WorkHandler))
	log.Fatal(http.ListenAndServe("localhost:"+PORT, nil))

}

type Request struct {
	Id     int    `json:"id"`
	Action string `json:"action"`
}

func WorkHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println("received request. Doing some work...")

    // Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	request := Request{}
	_ = json.Unmarshal(reqBody, &request)

    // Do some work
	request.Id = request.Id + 1
	returnPayload, _ := json.Marshal(request)

    // Return response (+ serialize JSON payload)
	fmt.Fprint(w, string(returnPayload))
	// log.Println("done with the work.")
}
