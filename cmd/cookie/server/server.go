/*
#!/bin/bash

export CERT_DIR=/etc/certs/example.com
export APP_HOST=example.com;
export APP_PORT=8888 ;
export APP_ORGANIZATION=${APP_HOST} ;
export APP_CERT=${CERT_DIR}/example.com.crt ;
export APP_KEY=${CERT_DIR}/example.com.key ;
export APP_PATH=dist;

# local variables:
# mode: shell-script
# end:

*/
// https://gowebexamples.com/sessions/
/*
$ go run sessions.go

$ curl -s http://localhost:8080/secret
Forbidden

$ curl -s -I http://localhost:8080/login
Set-Cookie: cookie-name=MTQ4NzE5Mz...

$ curl -s --cookie "cookie-name=MTQ4NzE5Mz..." http://localhost:8080/secret
The cake is a lie!
*/
// sessions.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/sessions"

	"github.com/davidwalter0/go-cfg"
	// "github.com/davidwalter0/go-tracer"
	"github.com/davidwalter0/tools/trace/httptrace"
	"github.com/jehiah/go-strftime"
)

var (
	app  App
	done = make(chan bool)

	// Build version build string
	Build string

	// Commit version commit string
	Commit string
)

// Now current formatted time
func Now() string {
	format := "%Y.%m.%d.%H.%M.%S.%z"
	now := time.Now()
	return strftime.Format(format, now)
}

func init() {
	var err error
	if err = cfg.ProcessHoldFlags("APP", &app); err != nil {
		log.Fatalf("%v\n", err)
	}
	cfg.Freeze()
	array := strings.Split(os.Args[0], "/")
	if false {
		me := array[len(array)-1]
		fmt.Println(me, "version built at:", Build, "commit:", Commit)
	}
}

// App application configuration struct
type App struct {
	Cert string
	Key  string
	Host string
	Port string
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	// dump(w, r)
	session, _ := store.Get(r, "vpn0.me")
	fmt.Printf("%v\n", session)
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	// dump(w, r)

	// dump(w, r)
	session, _ := store.Get(r, "vpn0.me")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	// dump(w, r)
	session, _ := store.Get(r, "vpn0.me")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	var err error
	http.HandleFunc("/secret", httptrace.EchoMeta(secret))
	http.HandleFunc("/login", httptrace.EchoMeta(login))
	http.HandleFunc("/logout", httptrace.EchoMeta(logout))

	fmt.Println("PORT on which  " + ":" + app.Port)
	fmt.Println("HOST interface " + ":" + app.Host)
	fmt.Printf("HTTPS/Listening on %s:%s\n", app.Host, app.Port)
	var url = fmt.Sprintf("%s:%s", app.Host, app.Port)
	if err = http.ListenAndServeTLS(url, app.Cert, app.Key, nil); err != nil {
		log.Fatal(url, err)
	}
}
