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

func TestDoRecordRequestInvalidRequest(t *testing.T) {
	_, err := doRecordRequest(&http.Request{})
	result := err.Error()

	expected := "error in Do"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}
}

func TestDoRecordRequestIOError(t *testing.T) {
	originalClient := httpClient
	httpClient = &mockedClient{}

	_, err := doRecordRequest(&http.Request{})
	result := err.Error()

	expected := "error in ReadAll"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}

	httpClient = originalClient
}

func TestDoRecordRequestStatusNotOK(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	method := "GET"
	hogeURL := "http://hoge.com"
	mockStr := "as a mock"
	badStatus := http.StatusBadRequest

	httpmock.RegisterResponder(method, hogeURL, httpmock.NewStringResponder(badStatus, mockStr))
	req, _ := http.NewRequest(method, hogeURL, nil)

	_, err := doRecordRequest(req)
	result := errors.Cause(err).Error()

	expected := fmt.Sprintf("error body: %s", mockStr)
	if result != expected {
		t.Errorf("expected '%s', bug got '%s'", expected, result)
	}
}

func TestDoRecordRequestBadJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	method := "GET"
	hogeURL := "http://hoge.com"
	badJSON := "{hoge}"

	httpmock.RegisterResponder(method, hogeURL, httpmock.NewStringResponder(http.StatusOK, badJSON))
	req, _ := http.NewRequest(method, hogeURL, nil)

	_, err := doRecordRequest(req)
	result := err.Error()

	expected := "error in Decode"
	if strings.Index(result, expected) != 0 {
		t.Errorf("expected '%s', bug got '%s'", expected, result)
	}
}

func TestDoRecordRequestValidResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	method := "GET"
	hogeURL := "http://hoge.com"

	expected := `{"record":[{"id":"hoge","name":"fuga","type":"A","prio":"10","content":"192.168.0.1","ttl":"10"}]}`
	httpmock.RegisterResponder(method, hogeURL, httpmock.NewStringResponder(http.StatusOK, expected))
	req, _ := http.NewRequest(method, hogeURL, nil)

	resultResp, _ := doRecordRequest(req)
	result, err := json.Marshal(&resultResp)
	if err != nil {
		t.Errorf("error in Marshal: %v", err)
		return
	}

	if string(result) != expected {
		t.Errorf("expected '%+v', bug got '%+v'", expected, string(result))
	}
}
func TestDoRecordRequestEmptyResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	method := "GET"
	hogeURL := "http://hoge.com"

	emptyResp := `[]`
	httpmock.RegisterResponder(method, hogeURL, httpmock.NewStringResponder(http.StatusOK, emptyResp))
	req, _ := http.NewRequest(method, hogeURL, nil)

	resultResp, _ := doRecordRequest(req)
	result, err := json.Marshal(&resultResp)
	if err != nil {
		t.Errorf("error in Marshal: %v", err)
		return
	}

	expected := `{"record":[]}`
	if string(result) != expected {
		t.Errorf("expected '%+v', bug got '%+v'", expected, string(result))
	}
}
