package httptrace

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// FormatHeaderAsRequestTable return as string
func FormatHeaderAsRequestTable(k, v interface{}, r *http.Request) (text string) {
	switch v.(type) {
	case string:
		text += fmt.Sprintf("%-32.32s -%48.48s\n", k, v)
	case []string:
		var cookies = v.([]string)
		if len(cookies) == 1 {
			text += fmt.Sprintf("%-32.32s %-48.48s\n", k, cookies[0])
		} else {
			for _, cookie := range cookies {
				text += fmt.Sprintf("%-32.32s %-48.48s\n", " ", cookie)
			}
		}
	}
	return
}

// FormatRequestTable format request info
func FormatRequestTable(r *http.Request) (text string) {
	if len(r.Header) > 0 {
		text += fmt.Sprintf("Header Fields\n")
		text += fmt.Sprintf("%-32.32s %-32.32s\n", "Field", "Value")
		for k, v := range r.Header {
			text += FormatHeaderAsRequestTable(k, v, r)
		}
	}
	return
}

// RequestMeta for request returned as string
func RequestMeta(r *http.Request) (text string) {
	text += fmt.Sprintf("%-32.32s %-32.32s %-32.32s\n", "Method", "Protocol", "URL")
	text += fmt.Sprintf("%-32.32s %-32.32s %-32.32s \n", r.Method, r.URL, r.Proto)
	text += fmt.Sprintf("%-32.32s %-32.32s\n", "Host", "Address")
	text += fmt.Sprintf("%-32.32s %-32.32s\n", r.Host, r.RemoteAddr)
	// Get value for a specified token
	text += fmt.Sprintf("\n\nFinding value of \"Accept\" %q\n\n", r.Header["Accept"])
	//Iterate over all header fields
	text += fmt.Sprintf("%s\n", FormatRequestTable(r))
	var body, _ = ioutil.ReadAll(r.Body)
	r.Body.Close()
	text += fmt.Sprintf("Body\n%s\n\n", string(body))
	return
}

// DumpRequestMeta to fmt package stdout
func DumpRequestMeta(w http.ResponseWriter, r *http.Request) {
	fmt.Println(RequestMeta(r))
}

// EchoRequestMeta collect and wcho request metadata and body back to
// response writer http.ResponseWriter
func EchoRequestMeta(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, RequestMeta(r))
}

// ChainableEchoRequestMeta schedule close of body after completion
func ChainableEchoRequestMeta(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		EchoRequestMeta(w, r)
		next(w, r)
	}
}

// EchoMeta echo header and info schedule close of body after completion
func EchoMeta(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		EchoRequestMeta(w, r)
		next(w, r)
	}
}
