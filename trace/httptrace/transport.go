package httptrace // import "github.com/davidwalter0/tools/trace/httptrace"

import (
	. "github.com/davidwalter0/tools/util"
	"github.com/davidwalter0/tools/util/tlsconfig"

	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptrace"
)

// TransportTracer is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type TransportTracer struct {
	current *http.Request
	http.Transport
}

// NewTLSClientTracedRequest from URI and request action type in
// GET,POST,PUT,HEADER...
func NewTLSClientTracedRequest(Cert, Key, Ca, URI, Action string) (*http.Client, *http.Request, error) {
	var req *http.Request
	var err error
	var client *http.Client
	var trTrace = &TransportTracer{}
	var trace *httptrace.ClientTrace
	var tlsConfig *tls.Config
	// var transport *http.Transport
	//	var transport *TransportTracer

	if req, err = http.NewRequest(Action, URI, nil); err != nil {
		return nil, nil, err
	}

	trace = trTrace.NewClientTrace()

	// transport = &TransportTracer{
	// 	current:         req,
	// 	TLSClientConfig: tlsConfig,
	// }

	// req = req.WithContext(httptrace.WithClientTrace(req.Context(), transport))

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	// transport = &http.Transport{
	// 	current:         req,
	// 	TLSClientConfig: tlsConfig,
	// }

	if tlsConfig, err = tlsconfig.NewTLSConfig(Cert, Key, Ca); err != nil {
		return nil, nil, err
	}
	client = &http.Client{
		Transport: &TransportTracer{
			current:   req,
			Transport: http.Transport{TLSClientConfig: tlsConfig},
		},
		// Transport: &http.Transport{
		// 	//			current:         req,
		// 	TLSClientConfig: tlsConfig,
		// },
	}

	return client, req, nil
}

// NewClientTracedRequest
func NewClientTracedRequest(req *http.Request) (*http.Client, *http.Request) {
	var trTrace *TransportTracer
	trTrace, req = NewTransportTracerRequest(req)
	// **FIXME**	return &http.Client{}, req
	return &http.Client{Transport: trTrace}, req
}

// NewTransportTracerRequest
// **FIXME**
// func NewTransportTracerRequest(req *http.Request) (*http.Request) {
func NewTransportTracerRequest(req *http.Request) (*TransportTracer, *http.Request) {
	var trTrace = &TransportTracer{}
	trace := trTrace.NewClientTrace()
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	// **FIXME**
	// return req
	return trTrace, req
}

// NewClientTrace
func (trace *TransportTracer) NewClientTrace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GetConn:           trace.GetConn,
		GotConn:           trace.GotConn,
		DNSStart:          trace.DNSStart,
		DNSDone:           trace.DNSDone,
		TLSHandshakeStart: trace.TLSHandshakeStart,
		TLSHandshakeDone:  trace.TLSHandshakeDone,
	}
}

// RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (trace *TransportTracer) RoundTrip(req *http.Request) (*http.Response, error) {
	trace.current = req
	return http.DefaultTransport.RoundTrip(req)
}

// GotConn prints whether the connection has been used previously
// for the current request.
func (trace *TransportTracer) GotConn(info httptrace.GotConnInfo) {
	if trace.current != nil {
		fmt.Printf("GotConn for %v reused? %v\n", trace.current.URL, info.Reused)
	}
}

// GetConn prints whether the connection has been used previously
// for the current request.
func (trace *TransportTracer) GetConn(hostPort string) {
	fmt.Printf("GetConn hostPort %s\n", hostPort)
}

func (trace *TransportTracer) DNSDone(dnsInfo httptrace.DNSDoneInfo) {
	fmt.Printf("DNSDone %v\n", JSONify(dnsInfo))
}

func (trace *TransportTracer) DNSStart(dnsInfo httptrace.DNSStartInfo) {
	fmt.Printf("DNSStart %v\n", JSONify(dnsInfo))
}

func (trace *TransportTracer) TLSHandshakeStart() {
	if trace.current != nil {
		fmt.Printf("TLSHandshakeStart %v\n", trace.current.URL)
	} else {
		fmt.Printf("TLSHandshakeStart\n")
	}
}

func (trace *TransportTracer) TLSHandshakeDone(state tls.ConnectionState, err error) {
	if trace.current != nil {
		fmt.Printf("TLSHandshakeDone %v\nState: %s\n%v\n", trace.current.URL, JSONify(state), err)
	} else {
		fmt.Printf("TLSHandshakeDone State: %s\n%v\n", JSONify(state), err)
	}
}
