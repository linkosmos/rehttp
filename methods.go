package rehttp

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Request methods
const (
	GET    = "GET"
	HEAD   = "HEAD"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

// Do - sends an HTTP request and returns an HTTP response, following
// policy (e.g. redirects, cookies, auth) as configured on the client.
//
// An error is returned if caused by client policy (such as
// CheckRedirect), or if there was an HTTP protocol error.
// A non-2xx response doesn't cause an error.
//
// When err is nil, resp always contains a non-nil resp.Body.
//
// Callers should close resp.Body when done reading from it. If
// resp.Body is not closed, the Client's underlying RoundTripper
// (typically Transport) may not be able to re-use a persistent TCP
// connection to the server for a subsequent "keep-alive" request.
//
// The request Body, if non-nil, will be closed by the underlying
// Transport, even on errors.
//
// Generally Get, Post, or PostForm will be used instead of Do.
func (r *Client) Do(req *http.Request) (*http.Response, error) {
	return r.Client.Do(req)
}

// RoundTrip implements the RoundTripper interface.
//
// For higher-level HTTP client support (such as handling of cookies
// and redirects), see Get, Post, and the Client type.
func (r Client) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.Transport.RoundTrip(req)
}

// GET - request
func (r *Client) GET(u *url.URL) *http.Request {
	return r.NewRequest(GET, u, nil)
}

// HEAD - request
func (r *Client) HEAD(u *url.URL) *http.Request {
	return r.NewRequest(HEAD, u, nil)
}

// NewRequest returns a new Request given a method, URL, and optional body.
//
// If the provided body is also an io.Closer, the returned
// Request.Body is set to body and will be closed by the Client
// methods Do, Post, and PostForm, and Transport.RoundTrip.
func (r *Client) NewRequest(method string, u *url.URL, body io.Reader) (req *http.Request) {
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}
	req = &http.Request{
		Method:     method,
		URL:        u,
		Proto:      r.RequestProto,
		ProtoMajor: r.RequestProtoMajor,
		ProtoMinor: r.RequestProtoMinor,
		Header:     r.ClientOptions.Headers,
		Body:       rc,
		Host:       u.Host,
	}
	if body != nil {
		switch v := body.(type) {
		case *bytes.Buffer:
			req.ContentLength = int64(v.Len())
		case *bytes.Reader:
			req.ContentLength = int64(v.Len())
		case *strings.Reader:
			req.ContentLength = int64(v.Len())
		}
	}
	return req
}
