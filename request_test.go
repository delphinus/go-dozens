package dozens

import (
	"net/url"
	"testing"

	"github.com/delphinus/go-dozens/endpoint"
)

func TestMakeGet(t *testing.T) {
	hogeURL, _ := url.Parse("http://hoge.com")
	p := endpoint.Endpoint{
		Base:  hogeURL,
		Chunk: "",
	}
	if _, err := MakeGet("", p); err != nil {
		t.Errorf("MakeGet returned error: %v", err)
	}
}

func TestMakePost(t *testing.T) {
	hogeURL, _ := url.Parse("http://hoge.com")
	p := endpoint.Endpoint{
		Base:  hogeURL,
		Chunk: "",
	}
	if _, err := MakePost("", p, nil); err != nil {
		t.Errorf("MakeGet returned error: %v", err)
	}
}

func TestMakeDelete(t *testing.T) {
	hogeURL, _ := url.Parse("http://hoge.com")
	p := endpoint.Endpoint{
		Base:  hogeURL,
		Chunk: "",
	}
	if _, err := MakeDelete("", p); err != nil {
		t.Errorf("MakeGet returned error: %v", err)
	}
}

func TestMakeGetWithError(t *testing.T) {
	originalMethodGet := methodGet
	methodGet = "(" // invalid method rune

	p := endpoint.Endpoint{
		Base:  &url.URL{},
		Chunk: "",
	}
	if _, err := MakeGet("", p); err == nil {
		t.Errorf("MakeGet did not return error")
	}

	methodGet = originalMethodGet
}
