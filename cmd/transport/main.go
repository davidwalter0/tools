package main

import (
	. "../transport"

	"log"
	"net/http"
)

func main() {
	var client *http.Client
	var req *http.Request
	// req, _ = http.NewRequest("GET", "https://google.com", nil)
	req, _ = http.NewRequest("GET", "https://vpn0.me:5556", nil)
	client, req = NewClientTracedRequest(req)
	if _, err := client.Do(req); err != nil {
		log.Fatal(err)
	}
}
