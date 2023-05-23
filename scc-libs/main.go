package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func WorkHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("received request. Doing some work...")

    // Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	request := Request{}
	_ = json.Unmarshal(reqBody, &request)

    // Do some work
	request.Id = request.Id + 1
	returnPayload, _ := json.Marshal(request)

    // Return response (+ serialize JSON payload)
	fmt.Fprint(w, string(returnPayload))
	log.Println("done with the work.")
}
