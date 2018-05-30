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
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/davidwalter0/tools/trace/httptrace"
)

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
	http.HandleFunc("/secret", httptrace.EchoMeta(secret))
	http.HandleFunc("/login", httptrace.EchoMeta(login))
	http.HandleFunc("/logout", httptrace.EchoMeta(logout))
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

// func dump(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(trace.RequestMeta(r))
// }

// func dump(w http.ResponseWriter, r *http.Request) {
// 	for k, v := range r.Header {
// 		// fmt.Printf("%-32.32s -%48.48s\n", k,v)
// 		var typeT = interface{}(v)
// 		switch typeT.(type) {
// 		case string:
// 			fmt.Printf("%-32.32s -%48.48s\n", k, v)
// 		case []string:
// 			fmt.Printf("%-32.32s\n", k)
// 			for x, cookie := range v {
// 				fmt.Printf("%4d %s", x, cookie)
// 			}
// 		}
// 		fmt.Println()
// 	}

// }
