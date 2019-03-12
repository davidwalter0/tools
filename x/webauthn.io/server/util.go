package server

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

func Transport(tlsConfig *tls.Config) *http.Transport {
	return &http.Transport{
		TLSClientConfig: tlsConfig,
	}
}

func TLSConfig(Ca string) (*tls.Config, error) {
	var err error
	var caCert []byte
	var caCertAuthPool *x509.CertPool
	var tlsConfig *tls.Config

	// Load CA cert
	if caCert, err = ioutil.ReadFile(Ca); err != nil {
		return nil, err
	}

	caCertAuthPool = x509.NewCertPool()
	caCertAuthPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
		//		Certificates: []tls.Certificate{cert},
		RootCAs: caCertAuthPool,
	}

	tlsConfig.BuildNameToCertificate()
	return tlsConfig, nil
}
