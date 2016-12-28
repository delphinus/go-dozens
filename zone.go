package dozens

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// ZoneResponse means response of zones
type ZoneResponse struct {
	Domain []domain `json:"domain"`
}

type domain struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func doZoneRequest(req *http.Request) (ZoneResponse, error) {
	zoneResp := ZoneResponse{}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in Do")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return zoneResp, errors.Wrap(err, "error in ReadAll")
		}
		return zoneResp, errors.Errorf("error body: %s", body)
	}

	if err := json.NewDecoder(resp.Body).Decode(&zoneResp); err != nil {
		return zoneResp, errors.Wrap(err, "error in Decode")
	}

	return zoneResp, nil
}
