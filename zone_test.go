package dozens

import (
	"fmt"
	"net/http"
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

	httpmock.RegisterResponder(method, hogeURL, httpmock.NewStringResponder(400, mockStr))
	req, _ := http.NewRequest(method, hogeURL, nil)
	expected := fmt.Sprintf("error body: %s", mockStr)

	_, err := doZoneRequest(req)
	result := errors.Cause(err).Error()

	if result != expected {
		t.Errorf("expected '%s', bug got '%s'", expected, result)
	}
}
