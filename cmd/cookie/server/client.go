package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/davidwalter0/go-cfg"
	// "github.com/davidwalter0/go-tracer"
	//"github.com/davidwalter0/tools/trace/httptrace"
)

type App struct {
	Host string `json:"host" doc:"Host for domain info"`
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

var SetCookie = "Set-Cookie"

func main() {
	const (
		login  = "/login"
		secret = "/secret"
	)
	var rsp *http.Response

	var err error
	var client *http.Client
	var tlsConfig *tls.Config
	// var request *http.Request
	//NewTLSClientTracedRequest(Cert, Key, Ca, URI, Action string)
	// client, request, err = httptrace.NewTLSClientTracedRequest(
	// 	// client, _, err = httptrace.NewTLSClientTracedRequest(
	// 	app.Cert,
	// 	app.Key,
	// 	app.Ca,
	// 	app.URI+login,
	// 	"GET",
	// )
	var transport *http.Transport
	if tlsConfig, err = TLSConfig(app.Ca); err != nil {
		log.Fatal(err)
	}
	transport = Transport(tlsConfig)
	// client := &http.Client{
	// 	Jar: jar,

	// }

	client = &http.Client{
		Transport: transport,
		Jar:       CookieJar(),
	}

	postData := url.Values{}
	//	postData.Set("keyword", "尹相杰")
	postData.Set("keyword", "KeywordValue")

	var r *http.Request
	if r, err = http.NewRequest("POST", app.URI+login, strings.NewReader(postData.Encode())); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r.Header.Add("Content-Type", "application/json")
	if rsp, err = client.Do(r); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var data []byte
	// Dump response
	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Login", string(data))
	r.Body.Close()

	if r, err = http.NewRequest("POST", app.URI+secret, strings.NewReader(postData.Encode())); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r.Header.Add("Content-Type", "application/json")
	// r.Header.Add("Cookie", fmt.Sprintf("%s=%s", SetCookie, cookieValue))
	// r.Header.Add("Cookie", fmt.Sprintf("%s=%s", MapKey2, cookieValue))
	if rsp, err = client.Do(r); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rsp.Body.Close()

	// Dump response
	data, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Secret", string(data))
}

func CookieJar() *cookiejar.Jar {
	var login = "/login"
	var MapKey = "Set-Cookie"
	jar, _ := cookiejar.New(nil)
	var cookies = []*http.Cookie{}
	var cookieValue = "MTUyNzQ3OTE1MnxEdi1CQkFFQ180SUFBUkFCRUFBQUpmLUNBQUVHYzNSeWFXNW5EQThBRFdGMWRHaGxiblJwWTJGMFpXUUVZbTl2YkFJQ0FBRT18jT6qXznIa7VOltg_b0j9XQb8IVx2ABqp7JNkVs_QN5g="
	cookie := &http.Cookie{
		Name:   MapKey,
		Value:  cookieValue,
		Path:   "/",
		Domain: app.Host,
	}
	cookies = append(cookies, cookie)
	u, _ := url.Parse(app.URI + login)
	jar.SetCookies(u, cookies)
	fmt.Println(jar.Cookies(u))
	return jar
}
