package dozens

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestGetAuthorizeWithNewRequestError(t *testing.T) {
	originalMethodGet := methodGet
	methodGet = "(" // invalid method
	defer func() { methodGet = originalMethodGet }()

	_, err := GetAuthorize("", "")
	result := err.Error()

	expected := "error in NewRequest"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}
}

type mockedErrorClient struct{}

func (c *mockedErrorClient) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("hoge error")
}

func TestGetAuthorizeWithDoError(t *testing.T) {
	originalClient := httpClient
	httpClient = &mockedErrorClient{}
	defer func() { httpClient = originalClient }()

	_, err := GetAuthorize("", "")
	result := err.Error()

	expected := "error in Do"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}
}

func TestGetAuthorizeWithReadAllError(t *testing.T) {
	originalClient := httpClient
	httpClient = &mockedClient{}
	defer func() { httpClient = originalClient }()

	_, err := GetAuthorize("", "")
	result := err.Error()

	expected := "error in ReadAll"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}
}
