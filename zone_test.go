package dozens

import (
	"encoding/json"
	"fmt"
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

func TestDoZoneRequestStatusNotOK(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	method := "GET"
	hogeURL := "http://hoge.com"
	mockStr := "as a mock"
	badStatus := http.StatusBadRequest

	httpmock.RegisterResponder(method, hogeURL, httpmock.NewStringResponder(badStatus, mockStr))
	req, _ := http.NewRequest(method, hogeURL, nil)

	_, err := doZoneRequest(req)
	result := errors.Cause(err).Error()

	expected := fmt.Sprintf("error body: %s", mockStr)
	if result != expected {
		t.Errorf("expected '%s', bug got '%s'", expected, result)
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
		t.Errorf("expected '%s', bug got '%s'", expected, result)
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
		t.Errorf("expected '%+v', bug got '%+v'", expected, result)
	}
}
