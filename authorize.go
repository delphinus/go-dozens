// Package dozens is a library for accessing the service Dozens.
//
// Dozens ( http://dozens.jp ) is a DNS service that has a simple Web interface
// and high functionality.  It has published API Reference (
// http://help.dozens.jp/categories/apiリファレンス/ ) and can be managed from
// any CLI tools.
//
// This package dozens is an implementation for the whole API.  This has been
// fully tested and has much reliability.
package dozens

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/delphinus/go-dozens/endpoint"
	"github.com/pkg/errors"
)

// AuthorizeResponse means response of authorize.  It has the token string.
type AuthorizeResponse struct {
	AuthToken string `json:"auth_token"`
}

// GetAuthorize returns authorization info.
//   resp, err := dozens.GetAuthorize("some key", "some user")
//   if err != nil {
//     panic(err)
//   }
//
//   token := resp.AuthToken
func GetAuthorize(key, user string) (AuthorizeResponse, error) {
	authorizeResp := AuthorizeResponse{}

	req, err := http.NewRequest(methodGet, endpoint.Authorize().String(), nil)
	if err != nil {
		return authorizeResp, errors.Wrap(err, "error in NewRequest")
	}
	req.Header.Set("X-Auth-Key", key)
	req.Header.Set("X-Auth-User", user)

	resp, err := httpClient.Do(req)
	if err != nil {
		return authorizeResp, errors.Wrap(err, "error in Do")
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return authorizeResp, errors.Wrap(err, "error in ReadAll")
		}
		return authorizeResp, errors.Errorf("error body: %s", body)
	}

	if err := json.NewDecoder(resp.Body).Decode(&authorizeResp); err != nil {
		return authorizeResp, errors.Wrap(err, "error in Decode")
	}

	return authorizeResp, nil
}
