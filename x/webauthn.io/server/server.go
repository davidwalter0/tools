package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"log"

	"github.com/davidwalter0/tools/x/webauthn.io/config"
	"github.com/davidwalter0/tools/x/webauthn.io/session"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/gorilla/mux"
)

// Timeout is the number of seconds to attempt a graceful shutdown, or
// for timing out read/write operations
const Timeout = 5 * time.Second

// Option is an option that sets a particular value for the server
type Option func(*Server)

// Server is a configurable HTTP server that implements a demo of the WebAuthn
// specification
type Server struct {
	server   *http.Server
	config   *config.Config
	webauthn *webauthn.WebAuthn
	store    *session.Store
}

// NewServer returns a new instance of a Server configured with the provided
// configuration
func NewServer(config *config.Config, opts ...Option) (ws *Server, err error) {
	var tlsConfig *tls.Config

	if tlsConfig, err = TLSConfig(config.Ca); err != nil {
		log.Fatal(err)
	}

	// addr := net.JoinHostPort("vpn0.me", config.HostPort)
	addr := net.JoinHostPort(config.HostAddress, config.HostPort)
	defaultServer := &http.Server{
		Addr:         addr,
		ReadTimeout:  Timeout,
		WriteTimeout: Timeout,
		TLSConfig:    tlsConfig,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	defaultStore, err := session.NewStore()
	if err != nil {
		return nil, err
	}
	defaultWebAuthn, _ := webauthn.New(&webauthn.Config{
		RPDisplayName: config.RelyingParty,
		RPID:          config.RelyingParty,
	})

	ws = &Server{
		config:   config,
		server:   defaultServer,
		store:    defaultStore,
		webauthn: defaultWebAuthn,
	}
	for _, opt := range opts {
		opt(ws)
	}
	ws.registerRoutes()
	return ws, nil
}

// WithWebAuthn sets the webauthn configuration for the server
func WithWebAuthn(w *webauthn.WebAuthn) Option {
	return func(ws *Server) {
		ws.webauthn = w
	}
}

// Start starts the underlying HTTP server
func (ws *Server) Start() error {
	log.Printf("Starting webauthn server at %s", ws.server.Addr)
	fmt.Printf("srv %+v\n", ws)
	go func() {
		var routerHTTP = mux.NewRouter()
		routerHTTP.Handle("/", http.HandlerFunc(redirectToHTTPS(ws)))
		var url = ws.httpURI(!secure)
		fmt.Println("http url", url)
		http.ListenAndServe(url, routerHTTP)
	}()

	return ws.server.ListenAndServeTLS(ws.config.Cert, ws.config.Key)
	// return ws.server.ListenAndServe()
}

// Shutdown attempts to gracefully shutdown the underlying HTTP server.
func (ws *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	return ws.server.Shutdown(ctx)
}

func (ws *Server) registerRoutes() {
	router := mux.NewRouter()
	// Unauthenticated handlers for registering a new credential and logging in.
	router.HandleFunc("/", ws.Login)
	router.HandleFunc("/makeCredential/{name}", ws.RequestNewCredential).Methods("GET")
	router.HandleFunc("/makeCredential", ws.MakeNewCredential).Methods("POST")
	router.HandleFunc("/assertion/{name}", ws.GetAssertion).Methods("GET")
	router.HandleFunc("/assertion", ws.MakeAssertion).Methods("POST")
	router.HandleFunc("/user/{name}/exists", ws.UserExists).Methods("GET")

	// Authenticated handlers for viewing credentials after logging in
	router.HandleFunc("/dashboard", ws.LoginRequired(ws.Index))

	// Static file serving
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	ws.server.Handler = router
}

const secure = true

func (ws *Server) httpURI(secure bool) string {
	// return fmt.Sprintf("%s:%s", ws.config.HostAddress, func() string {
	if secure {
		fmt.Printf("secure %v port %s\n", secure, ws.config.HostPort)
		return fmt.Sprintf("%s:%s", ws.config.RelyingParty, ws.config.HostPort)
	}
	fmt.Printf("secure %v port %s\n", secure, ws.config.RedirectPort)
	return fmt.Sprintf("%s:%s", ws.config.HostAddress, ws.config.RedirectPort)
}

func redirectToHTTPS(ws *Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Redirect the incoming HTTP request. Note that "127.0.0.1:8081"
		// will only work if you are accessing the server from your local
		// machine.
		var url = fmt.Sprintf("https://%s", ws.httpsURI())
		log.Println("URL", url)
		http.Redirect(w, r, url+r.RequestURI, http.StatusMovedPermanently)
	}
}

func (ws *Server) httpsURI() string {
	return ws.httpURI(secure)
}
