package dozens

import (
	"io"
	"net/http"

	"github.com/delphinus/go-dozens/endpoint"
	"github.com/pkg/errors"
)

// MakeGet returns request for dozens
func MakeGet(token string, p endpoint.Endpoint) (*http.Request, error) {
	return makeRequest("GET", token, p, nil)
}

// MakePost returns request for dozens
func MakePost(token string, p endpoint.Endpoint, body io.Reader) (*http.Request, error) {
	return makeRequest("POST", token, p, body)
}

// MakeDelete returns request for dozens
func MakeDelete(token string, p endpoint.Endpoint) (*http.Request, error) {
	return makeRequest("DELETE", token, p, nil)
}

func makeRequest(method, token string, p endpoint.Endpoint, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, p.String(), body)
	if err != nil {
		return nil, errors.Wrap(err, "error in NewRequest")
	}

	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
