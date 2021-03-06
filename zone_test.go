package dozens

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/pkg/errors"
)

func TestDoZoneRequestInvalidRequest(t *testing.T) {
	_, err := doZoneRequest(&http.Request{})
	result := err.Error()

	expected := "error in Do"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}
}

type mockedBody struct {
	io.ReadCloser
}

func (b *mockedBody) Read(bytes []byte) (int, error) { return 0, errors.New("some error") }
func (b *mockedBody) Close() error                   { return nil }

type mockedClient struct{}

func (c *mockedClient) Do(req *http.Request) (*http.Response, error) {
	resp := http.Response{}
	resp.Body = &mockedBody{}
	return &resp, nil
}

func TestDoZoneRequestIOError(t *testing.T) {
	originalClient := httpClient
	httpClient = &mockedClient{}
	defer func() { httpClient = originalClient }()

	_, err := doZoneRequest(&http.Request{})
	result := err.Error()

	expected := "error in ReadAll"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}
}

func TestDoZoneRequestStatusNotOK(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	hogeURL := "http://hoge.com"
	mockStr := "as a mock"
	badStatus := http.StatusBadRequest

	httpmock.RegisterResponder(methodGet, hogeURL, httpmock.NewStringResponder(badStatus, mockStr))
	req, _ := http.NewRequest(methodGet, hogeURL, nil)

	_, err := doZoneRequest(req)
	result := errors.Cause(err).Error()

	expected := fmt.Sprintf("error status: %d, body: %s", http.StatusBadRequest, mockStr)
	if result != expected {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}
}

func TestDoZoneRequestBadJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	method := "GET"
	hogeURL := "http://hoge.com"
	badJSON := "{hoge}"

	httpmock.RegisterResponder(method, hogeURL, httpmock.NewStringResponder(http.StatusOK, badJSON))
	req, _ := http.NewRequest(method, hogeURL, nil)

	_, err := doZoneRequest(req)
	result := err.Error()

	expected := "error in Decode"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}
}

func TestDoZoneRequestValidResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	method := "GET"
	hogeURL := "http://hoge.com"

	expected := `{"domain":[{"id":"hoge","name":"fuga"}]}`
	httpmock.RegisterResponder(method, hogeURL, httpmock.NewStringResponder(http.StatusOK, expected))
	req, _ := http.NewRequest(method, hogeURL, nil)

	resultResp, _ := doZoneRequest(req)
	result, err := json.Marshal(&resultResp)
	if err != nil {
		t.Errorf("error in Marshal: %v", err)
		return
	}

	if string(result) != expected {
		t.Errorf("expected '%+v', bug got '%+v'", expected, string(result))
	}
}
