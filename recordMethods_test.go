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

func TestRecordCreateWithError(t *testing.T) {
	originalMethodPost := methodPost
	methodPost = "(" // invalid method
	defer func() { methodPost = originalMethodPost }()

	_, err := RecordCreate("", RecordCreateBody{})

	expected := "error in MakePost"
	result := err.Error()
	if strings.Index(result, expected) != 0 {
		t.Errorf("error does not found: %s", result)
	}
}

func TestRecordCreateValidResponse(t *testing.T) {
	mock := dozensMock{
		Method:   methodPost,
		URL:      endpoint.RecordCreate().String(),
		Response: validRecordResponse,
	}

	_, err := mock.Do(func() (interface{}, error) {
		return RecordCreate("", RecordCreateBody{})
	})

	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestRecordUpdateWithError(t *testing.T) {
	originalMethodPost := methodPost
	methodPost = "(" // invalid method
	defer func() { methodPost = originalMethodPost }()

	_, err := RecordUpdate("", "", RecordUpdateBody{})

	expected := "error in MakePost"
	result := err.Error()
	if strings.Index(result, expected) != 0 {
		t.Errorf("error does not found: %s", result)
	}
}

func TestRecordUpdateValidResponse(t *testing.T) {
	mock := dozensMock{
		Method:   methodPost,
		URL:      endpoint.RecordUpdate("").String(),
		Response: validRecordResponse,
	}

	_, err := mock.Do(func() (interface{}, error) {
		return RecordUpdate("", "", RecordUpdateBody{})
	})

	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestRecordDeleteWithError(t *testing.T) {
	originalMethodDelete := methodDelete
	methodDelete = "(" // invalid method
	defer func() { methodDelete = originalMethodDelete }()

	_, err := RecordDelete("", "")

	expected := "error in MakeDelete"
	result := err.Error()
	if strings.Index(result, expected) != 0 {
		t.Errorf("error does not found: %s", result)
	}
}

func TestRecordDeleteValidResponse(t *testing.T) {
	mock := dozensMock{
		Method:   methodDelete,
		URL:      endpoint.RecordDelete("").String(),
		Response: validRecordResponse,
	}

	_, err := mock.Do(func() (interface{}, error) {
		return RecordDelete("", "")
	})

	if err != nil {
		t.Errorf("error: %v", err)
	}
}
