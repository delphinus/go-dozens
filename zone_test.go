package dozens

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/pkg/errors"
)

func TestDoZoneRequestStatusNotOK(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	method := "GET"
	hogeURL := "http://hoge.com"
	mockStr := "as a mock"
	badStatus := http.StatusBadRequest

	httpmock.RegisterResponder(method, hogeURL, httpmock.NewStringResponder(badStatus, mockStr))
	req, _ := http.NewRequest(method, hogeURL, nil)
	expected := fmt.Sprintf("error body: %s", mockStr)

	_, err := doZoneRequest(req)
	result := errors.Cause(err).Error()

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
	expected := "error in Decode"

	_, err := doZoneRequest(req)
	result := err.Error()

	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', bug got '%s'", expected, result)
	}
}