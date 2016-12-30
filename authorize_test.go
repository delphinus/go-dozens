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

	_, err := GetAuthorize("", "")
	result := err.Error()

	expected := "error in NewRequest"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}

	methodGet = originalMethodGet
}

type mockedErrorClient struct{}

func (c *mockedErrorClient) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("hoge error")
}

func TestGetAuthorizeWithDoError(t *testing.T) {
	originalClient := httpClient
	httpClient = &mockedErrorClient{}

	_, err := GetAuthorize("", "")
	result := err.Error()

	expected := "error in Do"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}

	httpClient = originalClient
}

func TestGetAuthorizeWithReadAllError(t *testing.T) {
	originalClient := httpClient
	httpClient = &mockedClient{}

	_, err := GetAuthorize("", "")
	result := err.Error()

	expected := "error in ReadAll"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}

	httpClient = originalClient
}
