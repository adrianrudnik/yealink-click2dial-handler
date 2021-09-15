package internal

import (
	"crypto/tls"
	"net/http"
)

func GetClient() *http.Client {
	// Allow self-signed certs
	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: t,
	}

	return client
}
