package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var PORT string
var WORKER_PORT string

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Please provide ports!")
	}

	PORT = os.Args[1]
	WORKER_PORT = os.Args[2]
	fmt.Printf("Starting main service. Listening on localhost:%s...\n", PORT)

	http.HandleFunc("/", http.HandlerFunc(handleRequest))
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Path
	if strings.HasPrefix(u, "/thick") {
		req_body, _ := json.Marshal(&wo)
		rdr := bytes.NewReader(req_body)
		resp, err := http.Post("http://localhost:"+WORKER_PORT+"/do-work", "application/json", rdr)
		if err != nil {
			log.Fatal(err)
		}
		_, _ = io.ReadAll(resp.Body)
	} else if strings.HasPrefix(u, "/long") {
		req_body, _ := json.Marshal(&wo2)
		rdr := bytes.NewReader(req_body)
		resp, err := http.Post("http://localhost:"+WORKER_PORT+"/do-work2", "application/json", rdr)
		if err != nil {
			log.Fatal(err)
		}
		_, _ = io.ReadAll(resp.Body)
	}
	fmt.Fprint(w, "Hello, World!")
}

func createJSONPayload(size int) (payload []byte) {
	letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for i := 0; i < size; i = i + 1 {
		rndNmbr := rand.Intn(26)
		letter := letters[rndNmbr]
		payload = append(payload, byte(letter))
	}
	return
}

type WorkObject struct {
	Id      int
	Action  string
	Message string
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

var msg = createJSONPayload(500)
var wo = WorkObject{Id: 123, Action: "REQUEST_IMAGE", Message: string(msg)}
var wo2 = WorkObject2{Id: 123, Action: "REQUEST_USER",
	Message1:  string(createJSONPayload(10)),
	Message2:  string(createJSONPayload(10)),
	Message3:  string(createJSONPayload(10)),
	Message4:  string(createJSONPayload(10)),
	Message5:  string(createJSONPayload(10)),
	Message6:  string(createJSONPayload(10)),
	Message7:  string(createJSONPayload(10)),
	Message8:  string(createJSONPayload(10)),
	Message9:  string(createJSONPayload(10)),
	Message10: string(createJSONPayload(10)),
	Message11: string(createJSONPayload(10)),
	Message12: string(createJSONPayload(10)),
	Message13: string(createJSONPayload(10)),
	Message14: string(createJSONPayload(10)),
	Message15: string(createJSONPayload(10)),
	Message16: string(createJSONPayload(10)),
	Message17: string(createJSONPayload(10)),
	Message18: string(createJSONPayload(10)),
	Message19: string(createJSONPayload(10)),
	Message20: string(createJSONPayload(10)),
}
