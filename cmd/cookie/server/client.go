package main

import (
	// "crypto/tls"
	// "crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/davidwalter0/go-cfg"
	// "github.com/davidwalter0/go-tracer"
	"github.com/davidwalter0/tools/trace/httptrace"
)

type App struct {
	Cert string `json:"cert" doc:"A PEM eoncoded certificate file"`
	Key  string `json:"key"  doc:"A PEM encoded private key file"`
	Ca   string `json:"ca"   doc:"A PEM eoncoded CA certificate file"`
	URI  string `json:"uri"  doc:"Server URI"`
}

var (
	app App

	done = make(chan bool)

	// Build version build string
	Build string

	// Commit version commit string
	Commit string
)

func init() {
	var err error
	if err = cfg.ProcessHoldFlags("APP", &app); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cfg.Freeze()
	array := strings.Split(os.Args[0], "/")
	me := array[len(array)-1]
	fmt.Println(me, "version built at:", Build, "commit:", Commit)
}

func main() {
	const (
		login  = "/login"
		secret = "/secret"
	)
	var r *http.Response

	var err error
	var client *http.Client
	var request *http.Request
	//NewTLSClientTracedRequest(Cert, Key, Ca, URI, Action string)
	client, request, err = httptrace.NewTLSClientTracedRequest(
		// client, _, err = httptrace.NewTLSClientTracedRequest(
		app.Cert,
		app.Key,
		app.Ca,
		app.URI+login,
		"GET",
	)
	// client := &http.Client{Transport: transport}

	// // Do GET something
	// if r, err = client.Get(app.URI + login); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// if r, err = client.Get(app.URI + login); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	if r, err = client.Do(request); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer r.Body.Close()

	var data []byte
	// Dump response
	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))

	r, err = client.Get(app.URI + secret)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer r.Body.Close()

	// Dump response
	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
