package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	// "github.com/gorilla/sessions"

	"github.com/jehiah/go-strftime"

	"github.com/davidwalter0/go-cfg"
	//	. "github.com/davidwalter0/tools/util"
	"github.com/davidwalter0/tools/util/signalhandler"
)

var (
	err  error
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

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// ..

var routerHTTPS = mux.NewRouter()
var routerHTTP = mux.NewRouter()

func httpURI() string {
	return fmt.Sprintf("%s:%s", app.Host, app.Port)
}

func httpsURI() string {
	var err error
	var port int
	if port, err = strconv.Atoi(app.Port); err != nil {
		log.Fatalf("%v\n", err)
	}
	return fmt.Sprintf("%s:%d", app.Host, port+1)
}

// NewShutdown(gs *GracefulShutdown) only log info shutting down
func NewShutdown() (gs *signalhandler.GracefulShutdown) {
	gs = signalhandler.NewGracefulShutdown("Secure TLS Server Example", nil)

	gs.Graceful = func() {
		log.Printf("\n\n*Secure: %s\n", gs.Message)
		log.Printf("%s: handling signal %v shuting down in\n", gs.Name, gs.Handled)
		for i := 3; i > -1; i-- {
			fmt.Printf("%d\n", i)
			time.Sleep(1 * time.Second)
		}
		fmt.Println()
	}

	return gs
}

func main() {
	var gs = NewShutdown()

	fmt.Println("PORT on which  " + ":" + app.Port)
	fmt.Println("HOST interface " + ":" + app.Host)
	fmt.Printf("HTTPS/Listening on %s\n", httpsURI())
	fmt.Printf("HTTP redirect %s\n", httpURI())

	go func() {
		var url = httpsURI()

		routerHTTPS.HandleFunc("/", indexPageHandler)
		routerHTTPS.HandleFunc("/internal", internalPageHandler)
		routerHTTPS.HandleFunc("/login", loginHandler).Methods("POST")
		routerHTTPS.HandleFunc("/logout", logoutHandler).Methods("POST")
		http.Handle("/", routerHTTPS)

		var tlsConfig *tls.Config

		if tlsConfig, err = TLSConfig(app.Ca); err != nil {
			log.Fatal(err)
		}
		srv := &http.Server{
			Addr:         url,
			Handler:      routerHTTPS,
			TLSConfig:    tlsConfig,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}
		log.Fatal(srv.ListenAndServeTLS(app.Cert, app.Key))
	}()

	go func() {
		routerHTTP.Handle("/", http.HandlerFunc(redirectToHttps))
		http.ListenAndServe(httpURI(), routerHTTP)
	}()
	gs.Wait()
}

const indexPage = `
 <h1>Login</h1>
 <form method="post" action="/login">
     <label for="name">User name</label>
     <input type="text" id="name" name="name">
     <label for="password">Password</label>
     <input type="password" id="password" name="password">
     <button type="submit">Login</button>
 </form>
 `

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	//	log.Printf("%v\n", JSONify(cookieHandler))
	fmt.Fprintf(response, indexPage)
}

const internalPage = `
 <h1>Internal</h1>
 <hr>
 <small>User: %s</small>
 <form method="post" action="/logout">
     <button type="submit">Logout</button>
 </form>
 `

func internalPageHandler(response http.ResponseWriter, request *http.Request) {
	//	log.Printf("%v\n", JSONify(cookieHandler))
	userName := getUserName(request)
	if userName != "" {
		fmt.Fprintf(response, internalPage, userName)
	} else {
		http.Redirect(response, request, "/", 302)
	}
}

func loginHandler(response http.ResponseWriter, request *http.Request) {
	//	log.Printf("%v\n", JSONify(cookieHandler))
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, response)
		redirectTarget = "/internal"
	}
	http.Redirect(response, request, redirectTarget, 302)
}

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	//	log.Printf("%v\n", JSONify(cookieHandler))

	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
		log.Printf("%v\n", cookie)
	}
}

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		//	log.Printf("%v\n", JSONify(cookieHandler))
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			log.Printf("%v\n", cookie)
			log.Printf("%v\n", cookieValue)
			userName = cookieValue["name"]
		}
	}
	return userName
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// MTUyNzk5NDUzMXxyS1BrYmppWEtIdDByVDUybjFPQnZzTHg5V01sdVdfSUhySGRZS3R0a1JqcVNOd3A3WUJ3aTVMRlpwdERoOG89fGMi9k6ABEY-Itq50cqrgDcHApFyCBmAr5DXb2x9kLVj

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:8081"
	// will only work if you are accessing the server from your local
	// machine.
	var url = fmt.Sprintf("https://%s", httpsURI())
	log.Println("URL", url)
	http.Redirect(w, r, url+r.RequestURI, http.StatusMovedPermanently)
}
