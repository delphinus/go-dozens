package dozens

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// RecordResponse means response of zones
type RecordResponse struct {
	Record []record `json:"record"`
}

type record struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Prio    string `json:"prio"`
	Content string `json:"content"`
	TTL     string `json:"ttl"`
}

func doRecordRequest(req *http.Request) (RecordResponse, error) {
	recordResp := RecordResponse{[]record{}}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return recordResp, errors.Wrap(err, "error in Do")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recordResp, errors.Wrap(err, "error in ReadAll")
	}

	if resp.StatusCode != http.StatusOK {
		return recordResp, errors.Errorf("error body: %s", body)
	}

	// Dozens API has a bug. If records are nil, it responses `[]` instead of
	// `{"record":[]}`.
	if len(body) == 2 && string(body) == "[]" {
		return recordResp, nil
	}

	if err := json.Unmarshal(body, &recordResp); err != nil {
		return recordResp, errors.Wrap(err, "error in Decode")
	}

	return recordResp, nil
}
