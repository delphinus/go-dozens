package dozens

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/delphinus/go-dozens/endpoint"
	"github.com/jarcoal/httpmock"
)

type zoneMockOptions struct {
	Method    string
	URL       string
	Status    int
	DoRequest func() (ZoneResponse, error)
}

var validZoneResponse = ZoneResponse{
	Domain: []domain{
		domain{ID: "hoge", Name: "fuga"},
	},
}

func zoneMock(options zoneMockOptions) (ZoneResponse, error) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responder, _ := httpmock.NewJsonResponder(options.Status, validZoneResponse)
	httpmock.RegisterResponder(options.Method, options.URL, responder)

	return options.DoRequest()
}

func TestZoneCreateWithError(t *testing.T) {
	originalMethodPost := methodPost
	methodPost = "(" // invalid method
	defer func() { methodPost = originalMethodPost }()

	_, err := ZoneCreate("", ZoneCreateBody{})

	expected := fmt.Sprintf("error in MakePost")
	result := err.Error()
	if strings.Index(result, expected) != 0 {
		t.Errorf("error does not found: %s", result)
	}
}

func TestZoneCreateValidResponse(t *testing.T) {
	body := ZoneCreateBody{
		Name:            "hoge",
		AddGoogleApps:   true,
		GoogleAuthorize: "hoge",
		MailAddress:     "hoge",
	}
	_, err := zoneMock(zoneMockOptions{
		Method:    methodPost,
		URL:       endpoint.ZoneCreate().String(),
		Status:    http.StatusOK,
		DoRequest: func() (ZoneResponse, error) { return ZoneCreate("", body) },
	})

	if err != nil {
		t.Errorf("error: %v", err)
	}
}
