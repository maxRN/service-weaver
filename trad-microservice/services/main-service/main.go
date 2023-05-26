package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"strings"
)

const PORT = "8070"

var payload = createJSONPayload(500)
var requestBody = fmt.Sprintf("{\"id\": 123, \"action\": \"request_image\", \"message\": \"%s\"}", payload)

func main() {
	fmt.Printf("Starting main service. Listening on localhost:%s...\n", PORT)

	http.HandleFunc("/", http.HandlerFunc(handleRequest))
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Path
	if strings.HasPrefix(u, "/request") {
		sizeString := r.URL.Query().Get("size")
		size, _ := strconv.Atoi(sizeString)
		timesString := r.URL.Query().Get("times")
		times, _ := strconv.Atoi(timesString)
		doWork(size, times)
	}
	fmt.Fprint(w, "Hello, World!")
}

func doWork(size int, times int) {
	rdr := strings.NewReader(requestBody)
	resp, err := http.Post("http://localhost:8069/do-work", "application/json", rdr)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = io.ReadAll(resp.Body)
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
