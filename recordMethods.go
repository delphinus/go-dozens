package dozens

import (
	"bytes"
	"encoding/json"

	"github.com/delphinus/go-dozens/endpoint"
	"github.com/pkg/errors"
)

// RecordList returns records list
func RecordList(token, zoneName string) (RecordResponse, error) {
	recordResp := RecordResponse{}

	req, err := MakeGet(token, endpoint.RecordList(zoneName))
	if err != nil {
		return recordResp, errors.Wrap(err, "error in MakeGet")
	}

	return doRecordRequest(req)
}

// RecordCreateBody means post data for `create` request
type RecordCreateBody struct {
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Prio    string `json:"prio"`
	Content string `json:"content"`
	TTL     string `json:"ttl,omitempty"`
}

// RecordCreate creates record and returns records list
func RecordCreate(token string, body RecordCreateBody) (RecordResponse, error) {
	recordResp := RecordResponse{}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return recordResp, errors.Wrap(err, "error in Marshal")
	}

	req, err := MakePost(token, endpoint.RecordCreate(), bytes.NewBuffer(bodyJSON))
	if err != nil {
		return recordResp, errors.Wrap(err, "error in MakeGet")
	}

	return doRecordRequest(req)
}

// RecordDelete deletes record and returns records list
func RecordDelete(token, recordID string) (RecordResponse, error) {
	req, err := MakeDelete(token, endpoint.RecordDelete(recordID))
	if err != nil {
		return RecordResponse{}, errors.Wrap(err, "error in MakeDelete")
	}

	return doRecordRequest(req)
}

// RecordUpdateBody means post data for `update` request
type RecordUpdateBody struct {
	Prio    uint   `json:"prio,omitempty"`
	Content string `json:"content,omitempty"`
	TTL     string `json:"ttl,omitempty"`
}

// RecordUpdate updates record and returns records list
func RecordUpdate(token string, recordID string, body RecordUpdateBody) (RecordResponse, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return RecordResponse{}, errors.Wrap(err, "error in Marshal")
	}

	req, err := MakePost(token, endpoint.RecordUpdate(recordID), bytes.NewBuffer(bodyJSON))
	if err != nil {
		return RecordResponse{}, errors.Wrap(err, "error in MakeUpdate")
	}

	return doRecordRequest(req)
}
