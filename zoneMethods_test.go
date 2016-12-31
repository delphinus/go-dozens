package dozens

import (
	"net/http"
	"strings"
	"testing"

	"github.com/delphinus/go-dozens/endpoint"
	"github.com/jarcoal/httpmock"
)

type dozensMock struct {
	Method   string
	URL      string
	Status   int
	Response interface{}
}

func (m dozensMock) Do(reqFunc func() (interface{}, error)) (interface{}, error) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responder, _ := httpmock.NewJsonResponder(m.Status, m.Response)
	httpmock.RegisterResponder(m.Method, m.URL, responder)

	return reqFunc()
}

var validZoneResponse = ZoneResponse{
	Domain: []domain{
		domain{ID: "hoge", Name: "fuga"},
	},
}

func TestZoneListWithError(t *testing.T) {
	originalMethodGet := methodGet
	methodGet = "(" // invalid method
	defer func() { methodGet = originalMethodGet }()

	_, err := ZoneList("")

	expected := "error in MakeGet"
	result := err.Error()
	if strings.Index(result, expected) != 0 {
		t.Errorf("error does not found: %s", result)
	}
}

func TestZoneListValidResponse(t *testing.T) {
	mock := dozensMock{
		Method:   methodGet,
		URL:      endpoint.ZoneList().String(),
		Status:   http.StatusOK,
		Response: validZoneResponse,
	}

	_, err := mock.Do(func() (interface{}, error) {
		return ZoneList("")
	})

	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestZoneCreateWithError(t *testing.T) {
	originalMethodPost := methodPost
	methodPost = "(" // invalid method
	defer func() { methodPost = originalMethodPost }()

	_, err := ZoneCreate("", ZoneCreateBody{})

	expected := "error in MakePost"
	result := err.Error()
	if strings.Index(result, expected) != 0 {
		t.Errorf("error does not found: %s", result)
	}
}

func TestZoneCreateValidResponse(t *testing.T) {
	mock := dozensMock{
		Method:   methodPost,
		URL:      endpoint.ZoneCreate().String(),
		Status:   http.StatusOK,
		Response: validZoneResponse,
	}

	_, err := mock.Do(func() (interface{}, error) {
		return ZoneCreate("", ZoneCreateBody{})
	})

	if err != nil {
		t.Errorf("error: %v", err)
	}
}
