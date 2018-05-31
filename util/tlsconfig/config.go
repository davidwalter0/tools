package tlsconfig

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

// NewTLSConfig from URI and request action type in
// GET,POST,PUT,HEADER...
func NewTLSConfig(Cert, Key, Ca string) (*tls.Config, error) {
	var err error
	var cert tls.Certificate
	var caCert []byte
	var caCertAuthPool *x509.CertPool
	var tlsConfig *tls.Config

	// Load client cert
	if cert, err = tls.LoadX509KeyPair(Cert, Key); err != nil {
		return nil, err
	}

	// Load CA cert
	if caCert, err = ioutil.ReadFile(Ca); err != nil {
		return nil, err
	}

	caCertAuthPool = x509.NewCertPool()
	caCertAuthPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertAuthPool,
	}

	tlsConfig.BuildNameToCertificate()
	return tlsConfig, nil
}
