package dozens

import (
	"bytes"
	"encoding/json"

	"github.com/delphinus/go-dozens/endpoint"
	"github.com/pkg/errors"
)

// ZoneList returns zones list
func ZoneList(token string) (ZoneResponse, error) {
	zoneResp := ZoneResponse{}

	req, err := MakeGet(token, endpoint.ZoneList())
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in MakeGet")
	}

	return doZoneRequest(req)
}

// ZoneCreateBody means post data for `create` request
type ZoneCreateBody struct {
	Name            string `json:"name"`
	AddGoogleApps   bool   `json:"add_google_apps"`
	GoogleAuthorize string `json:"google_authorize,omitempty"`
	MailAddress     string `json:"mailaddress,omitempty"`
}

// ZoneCreate creates zone and returns zones list
func ZoneCreate(token string, body ZoneCreateBody) (ZoneResponse, error) {
	zoneResp := ZoneResponse{}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in Marshal")
	}

	req, err := MakePost(token, endpoint.ZoneCreate(), bytes.NewBuffer(bodyJSON))
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in MakeGet")
	}

	return doZoneRequest(req)
}

// ZoneUpdateBody means post data for `update` request
type ZoneUpdateBody struct {
	MailAddress string `json:"mailaddress"`
}

// ZoneUpdate creates zone and returns zones list
func ZoneUpdate(token, zoneID, mailAddress string) (ZoneResponse, error) {
	zoneResp := ZoneResponse{}
	body := ZoneUpdateBody{mailAddress}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in Marshal")
	}

	req, err := MakePost(token, endpoint.ZoneUpdate(zoneID), bytes.NewBuffer(bodyJSON))
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in MakeGet")
	}

	return doZoneRequest(req)
}

// ZoneDelete deletes zone and returns zones list
func ZoneDelete(token, zoneID string) (ZoneResponse, error) {
	req, err := MakeDelete(token, endpoint.ZoneDelete(zoneID))
	if err != nil {
		return ZoneResponse{}, errors.Wrap(err, "error in MakeGet")
	}

	return doZoneRequest(req)
}
