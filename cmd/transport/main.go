package main

import (
	"github.com/davidwalter0/tools/trace/httptrace"

	"log"
	"net/http"
)

func main() {
	var client *http.Client
	var req *http.Request
	// req, _ = http.NewRequest("GET", "https://google.com", nil)
	req, _ = http.NewRequest("GET", "https://vpn0.me:5556", nil)
	client, req = httptrace.NewClientTracedRequest(req)
	if _, err := client.Do(req); err != nil {
		log.Fatal(err)
	}
}
