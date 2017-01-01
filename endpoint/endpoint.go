package endpoint

import (
	"fmt"
	"net/url"
	"strings"
)

var baseURL *url.URL

func init() {
	url, _ := url.Parse("http://dozens.jp/api")
	baseURL = url
}

// Endpoint means the path struct
type Endpoint struct {
	Base  *url.URL
	Chunk string
}

// NewEndpoint returns Endpoint struct
func NewEndpoint(chunk string) Endpoint {
	return Endpoint{baseURL, chunk}
}

func (p Endpoint) String() string {
	u := url.URL{
		Scheme: p.Base.Scheme,
		Host:   p.Base.Host,
		Path:   strings.Join([]string{p.Base.Path, p.Chunk}, "/"),
	}
	return u.String()
}

// Authorize means `http://dozens.jp/api/authorize.json`
func Authorize() Endpoint {
	return NewEndpoint("authorize.json")
}

// ZoneList means `http://dozens.jp/api/zone.json`
func ZoneList() Endpoint {
	return NewEndpoint("zone.json")
}

// ZoneCreate means `http://dozens.jp/api/zone/create.json`
func ZoneCreate() Endpoint {
	return NewEndpoint("zone/create.json")
}

// ZoneUpdate means `http://dozens.jp/api/zone/update/:zone_id.json`
func ZoneUpdate(zoneID string) Endpoint {
	return NewEndpoint(fmt.Sprintf("zone/update/%s.json", zoneID))
}

// ZoneDelete means `http://dozens.jp/api/zone/delete/:zone_id.json`
func ZoneDelete(zoneID string) Endpoint {
	return NewEndpoint(fmt.Sprintf("zone/delete/%s.json", zoneID))
}

// RecordList means `http://dozens.jp/api/record/:zone_name.json`
func RecordList(zoneName string) Endpoint {
	return NewEndpoint(fmt.Sprintf("record/%s.json", zoneName))
}

// RecordCreate means `http://dozens.jp/api/record/create.json
func RecordCreate() Endpoint {
	return NewEndpoint("record/create.json")
}

// RecordDelete means `http://dozens.jp/api/record/delete/:record_id.json`
func RecordDelete(recordID string) Endpoint {
	return NewEndpoint(fmt.Sprintf("record/delete/%s.json", recordID))
}

// RecordUpdate means `http://dozens.jp/api/record/update/:record_id.json`
func RecordUpdate(recordID string) Endpoint {
	return NewEndpoint(fmt.Sprintf("record/update/%s.json", recordID))
}
