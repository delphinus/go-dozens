package dozens

import (
	"strings"
	"testing"

	"github.com/delphinus/go-dozens/endpoint"
)

var validRecordResponse = RecordResponse{
	Record: []record{
		record{
			ID:      "hoge",
			Name:    "hoge",
			Type:    "A",
			Prio:    "10",
			Content: "hoge",
			TTL:     "10",
		},
	},
}

func TestRecordListWithError(t *testing.T) {
	originalMethodGet := methodGet
	methodGet = "(" // invalid method
	defer func() { methodGet = originalMethodGet }()

	_, err := RecordList("", "")

	expected := "error in MakeGet"
	result := err.Error()
	if strings.Index(result, expected) != 0 {
		t.Errorf("error does not found: %s", result)
	}
}

func TestRecordListValidResponse(t *testing.T) {
	zone := "hoge"
	mock := dozensMock{
		Method:   methodGet,
		URL:      endpoint.RecordList(zone).String(),
		Response: validRecordResponse,
	}

	_, err := mock.Do(func() (interface{}, error) {
		return RecordList("", zone)
	})

	if err != nil {
		t.Errorf("error: %s", err)
	}
}
