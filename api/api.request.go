package api

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Request interface {
	Do(ctx context.Context, transport Transport) (*Response, error)
}

// newRequest creates an HTTP request.
func newRequest(method string, body io.Reader) (*http.Request, error) {

	r := http.Request{
		Method:     method,
		URL:        &url.URL{},
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}

	if body != nil {
		switch b := body.(type) {
		case *bytes.Buffer:
			r.Body = ioutil.NopCloser(body)
			r.ContentLength = int64(b.Len())
		case *bytes.Reader:
			r.Body = ioutil.NopCloser(body)
			r.ContentLength = int64(b.Len())
		case *strings.Reader:
			r.Body = ioutil.NopCloser(body)
			r.ContentLength = int64(b.Len())
		default:
			r.Body = ioutil.NopCloser(body)
		}
	}

	return &r, nil
}