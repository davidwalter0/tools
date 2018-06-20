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
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/sessions"

	"github.com/davidwalter0/go-cfg"
	// "github.com/davidwalter0/go-tracer"
	"github.com/davidwalter0/toolstrace/httptrace"

	"github.com/jehiah/go-strftime"
	// "github.com/davidwalter0/toolsutil/tlsconfig"
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
	Ca   string
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
	session, _ := store.Get(r, app.Host)
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
	session, _ := store.Get(r, app.Host)

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	// dump(w, r)
	session, _ := store.Get(r, app.Host)

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

type HandleFunc func(w http.ResponseWriter, req *http.Request)

func SecureHeaders(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.Write([]byte("This is an example server.\n"))
}

func Wrap(next HandleFunc, funcs ...HandleFunc) HandleFunc {
	if len(funcs) == 0 {
		return func(w http.ResponseWriter, r *http.Request) {
			next(w, r)
		}
	} else {
		return func(w http.ResponseWriter, r *http.Request) {
			next(w, r)
			Wrap(funcs[0], funcs[1:]...)(w, r)
		}
	}
}

func main() {
	// var err error
	// http.HandleFunc("/secret", httptrace.EchoMeta(secret))
	// http.HandleFunc("/login", httptrace.EchoMeta(login))
	// http.HandleFunc("/logout", httptrace.EchoMeta(logout))

	// fmt.Println("PORT on which  " + ":" + app.Port)
	// fmt.Println("HOST interface " + ":" + app.Host)
	// fmt.Printf("HTTPS/Listening on %s:%s\n", app.Host, app.Port)
	// var url = fmt.Sprintf("%s:%s", app.Host, app.Port)
	// if err = http.ListenAndServeTLS(url, app.Cert, app.Key, nil); err != nil {
	// 	log.Fatal(url, err)
	// }
	var err error
	// if err = http.ListenAndServeTLS(url, app.Cert, app.Key, nil); err != nil {
	// 	log.Fatal(url, err)
	// }
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is an example server.\n"))
	})

	mux.HandleFunc("/secret", Wrap(httptrace.EchoRequestMeta, SecureHeaders, secret))
	mux.HandleFunc("/login", Wrap(httptrace.EchoRequestMeta, SecureHeaders, login))
	mux.HandleFunc("/logout", Wrap(httptrace.EchoRequestMeta, SecureHeaders, logout))

	// mux.HandleFunc("/secret", httptrace.EchoMeta(secret))
	// mux.HandleFunc("/login", httptrace.EchoMeta(login))
	// mux.HandleFunc("/logout", httptrace.EchoMeta(logout))

	fmt.Println("PORT on which  " + ":" + app.Port)
	fmt.Println("HOST interface " + ":" + app.Host)
	fmt.Printf("HTTPS/Listening on %s:%s\n", app.Host, app.Port)
	var url = fmt.Sprintf("%s:%s", app.Host, app.Port)

	// cfg := &tls.Config{
	// 	MinVersion: tls.VersionTLS12,
	// 	// CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	// 	// PreferServerCipherSuites: true,
	// 	// CipherSuites: []uint16{
	// 	//     tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	// 	//     tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	// 	//     tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	// 	//     tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	// 	// },
	// }
	var tlsConfig *tls.Config

	// if tlsConfig, err = tlsconfig.NewTLSConfig(app.Cert, app.Key, app.Ca); err != nil {
	// 	log.Fatal(err)
	// }
	if tlsConfig, err = TLSConfig(app.Ca); err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Addr:         url,
		Handler:      mux,
		TLSConfig:    tlsConfig,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS(app.Cert, app.Key))
}
