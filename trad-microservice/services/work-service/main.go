package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

var PORT string

func main() {
    if len(os.Args) != 2 {
        log.Fatal("Please provide port number!")
    }
    PORT = os.Args[1]
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

    // Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	request := Request{}
	_ = json.Unmarshal(reqBody, &request)

    // Do some work
	request.Id = request.Id + 1
	returnPayload, _ := json.Marshal(request)

    // Return response (+ serialize JSON payload)
	fmt.Fprint(w, string(returnPayload))
}
