package rehttp

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/facebookgo/httpcontrol"
	"github.com/linkosmos/redial"
)

// -
const (
	RequestProto      = "HTTP/1.1"
	RequestProtoMinor = 0
	RequestProtoMajor = 1
)

// Client - http.: Transport, Client wrapper
type Client struct {
	RequestProto string

	RequestProtoMinor, RequestProtoMajor int

	// A Config structure is used to configure a TLS client or server.
	// After one has been passed to a TLS function it must not be
	// modified. A Config may be reused; the tls package will also not
	// modify it.
	TLS *tls.Config

	// Transport is an implementation of RoundTripper that supports HTTP,
	// HTTPS, and HTTP proxies (for either HTTP or HTTPS with CONNECT).
	// Transport can also cache connections for future re-use.
	Transport *httpcontrol.Transport

	// A Client is an HTTP client. Its zero value (DefaultClient) is a
	// usable client that uses DefaultTransport.
	//
	// The Client's Transport typically has internal state (cached TCP
	// connections), so Clients should be reused instead of created as
	// needed. Clients are safe for concurrent use by multiple goroutines.
	//
	// A Client is higher-level than a RoundTripper (such as Transport)
	// and additionally handles HTTP details such as cookies and
	// redirects.
	Client ClientRequester

	// rehttp.Client options
	ClientOptions *Options
}

// New - returns Request Client
func New(op *Options, address, port string) (r *Client, err error) {
	if op == nil {
		op = NewOptions()
	}
	dialer, err := redial.New(net.Dialer{
		Timeout:   op.DialerTimeout,
		Deadline:  op.DialerDeadline,
		KeepAlive: op.DialerKeepAlive,
		DualStack: op.DialerDualStack,
	}, address, port, op.ConnPoolInitial, op.ConnPoolMax)
	if err != nil {
		return nil, err
	}
	r = &Client{
		RequestProto:      RequestProto,
		RequestProtoMinor: RequestProtoMinor,
		RequestProtoMajor: RequestProtoMajor,
		TLS: &tls.Config{
			InsecureSkipVerify: op.TLSInsecureSkipVerify,
		},
		ClientOptions: op, // Setting rehttp.Client options
	}

	// Setting up TRANSPORT
	r.Transport = &httpcontrol.Transport{
		Dial:                  dialer.Dial,
		TLSClientConfig:       r.TLS,
		MaxTries:              op.TransportMaxTries,
		DisableKeepAlives:     op.TransportDisableKeepAlives,
		DisableCompression:    op.TransportDisableCompression,
		MaxIdleConnsPerHost:   op.TransportMaxIdleConnsPerHost,
		RequestTimeout:        op.TransportRequestTimeout,
		ResponseHeaderTimeout: op.TransportResponseHeaderTimeout,
	}

	// Setting up CLIENT, higher level API of TRANSPORT
	r.Client = &http.Client{
		Transport: r.Transport,
		Timeout:   op.ClientTimeout,
	}
	return r, nil
}
