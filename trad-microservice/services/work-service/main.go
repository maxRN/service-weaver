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
	http.Handle("/do-work2", http.HandlerFunc(WorkHandler2))
	log.Fatal(http.ListenAndServe("localhost:"+PORT, nil))

}

type WorkObject struct {
	Id      int    `json:"id"`
	Action  string `json:"action"`
	Message string `json:"message"`
}

func WorkHandler(w http.ResponseWriter, r *http.Request) {

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	request := WorkObject{}
	_ = json.Unmarshal(reqBody, &request)

	// Do some work
	request.Id = request.Id + 1
	returnPayload, _ := json.Marshal(request)

	// Return response (+ serialize JSON payload)
	fmt.Fprint(w, string(returnPayload))
}

func WorkHandler2(w http.ResponseWriter, r *http.Request) {

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	request := WorkObject2{}
	_ = json.Unmarshal(reqBody, &request)

	// Do some work
	request.Id = request.Id + 1
	returnPayload, _ := json.Marshal(request)

	// Return response (+ serialize JSON payload)
	fmt.Fprint(w, string(returnPayload))
}

type WorkObject2 struct {
	Id        int    `json:"id"`
	Action    string `json:"action"`
	Message1  string `json:"message1"`
	Message2  string `json:"message2"`
	Message3  string `json:"message3"`
	Message4  string `json:"message4"`
	Message5  string `json:"message5"`
	Message6  string `json:"message6"`
	Message7  string `json:"message7"`
	Message8  string `json:"message8"`
	Message9  string `json:"message9"`
	Message10 string `json:"message10"`
	Message11 string `json:"message11"`
	Message12 string `json:"message12"`
	Message13 string `json:"message13"`
	Message14 string `json:"message14"`
	Message15 string `json:"message15"`
	Message16 string `json:"message16"`
	Message17 string `json:"message17"`
	Message18 string `json:"message18"`
	Message19 string `json:"message19"`
	Message20 string `json:"message20"`
}
