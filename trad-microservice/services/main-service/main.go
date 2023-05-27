package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
)

var PORT string
var WORKER_PORT string

var payload = createJSONPayload(500)
var requestBody = fmt.Sprintf("{\"id\": 123, \"action\": \"request_image\", \"message\": \"%s\"}", payload)

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
	if strings.HasPrefix(u, "/request") {
		rdr := strings.NewReader(requestBody)
		resp, err := http.Post("http://localhost:"+WORKER_PORT+"/do-work", "application/json", rdr)
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
