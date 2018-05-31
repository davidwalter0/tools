package main

/*
type ClientTrace struct {
        // GetConn is called before a connection is created or
        // retrieved from an idle pool. The hostPort is the
        // "host:port" of the target or proxy. GetConn is called even
        // if there's already an idle cached connection available.
        GetConn func(hostPort string)

        // GotConn is called after a successful connection is
        // obtained. There is no hook for failure to obtain a
        // connection; instead, use the error from
        // Transport.RoundTrip.
        GotConn func(GotConnInfo)

        // PutIdleConn is called when the connection is returned to
        // the idle pool. If err is nil, the connection was
        // successfully returned to the idle pool. If err is non-nil,
        // it describes why not. PutIdleConn is not called if
        // connection reuse is disabled via Transport.DisableKeepAlives.
        // PutIdleConn is called before the caller's Response.Body.Close
        // call returns.
        // For HTTP/2, this hook is not currently used.
        PutIdleConn func(err error)

        // GotFirstResponseByte is called when the first byte of the response
        // headers is available.
        GotFirstResponseByte func()

        // Got100Continue is called if the server replies with a "100
        // Continue" response.
        Got100Continue func()

        // DNSStart is called when a DNS lookup begins.
        DNSStart func(DNSStartInfo)

        // DNSDone is called when a DNS lookup ends.
        DNSDone func(DNSDoneInfo)

        // ConnectStart is called when a new connection's Dial begins.
        // If net.Dialer.DualStack (IPv6 "Happy Eyeballs") support is
        // enabled, this may be called multiple times.
        ConnectStart func(network, addr string)

        // ConnectDone is called when a new connection's Dial
        // completes. The provided err indicates whether the
        // connection completedly successfully.
        // If net.Dialer.DualStack ("Happy Eyeballs") support is
        // enabled, this may be called multiple times.
        ConnectDone func(network, addr string, err error)

        // TLSHandshakeStart is called when the TLS handshake is started. When
        // connecting to a HTTPS site via a HTTP proxy, the handshake happens after
        // the CONNECT request is processed by the proxy.
        TLSHandshakeStart func()

        // TLSHandshakeDone is called after the TLS handshake with either the
        // successful handshake's connection state, or a non-nil error on handshake
        // failure.
        TLSHandshakeDone func(tls.ConnectionState, error)

        // WroteHeaders is called after the Transport has written
        // the request headers.
        WroteHeaders func()

        // Wait100Continue is called if the Request specified
        // "Expected: 100-continue" and the Transport has written the
        // request headers but is waiting for "100 Continue" from the
        // server before writing the request body.
        Wait100Continue func()

        // WroteRequest is called with the result of writing the
        // request and any body. It may be called multiple times
        // in the case of retried requests.
        WroteRequest func(WroteRequestInfo)
}

*/
