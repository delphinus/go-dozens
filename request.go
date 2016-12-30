package dozens

import (
	"io"
	"net/http"

	"github.com/delphinus/go-dozens/endpoint"
	"github.com/pkg/errors"
)

var methodGet = "GET"
var methodPost = "POST"
var methodDelete = "DELETE"

// MakeGet returns request for dozens
func MakeGet(token string, p endpoint.Endpoint) (*http.Request, error) {
	return makeRequest(methodGet, token, p, nil)
}

// MakePost returns request for dozens
func MakePost(token string, p endpoint.Endpoint, body io.Reader) (*http.Request, error) {
	return makeRequest(methodPost, token, p, body)
}

// MakeDelete returns request for dozens
func MakeDelete(token string, p endpoint.Endpoint) (*http.Request, error) {
	return makeRequest(methodDelete, token, p, nil)
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
