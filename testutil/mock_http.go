package testutil

import (
	"bytes"
	"io"
	"net/http"
)

type MockRoundTripper struct {
	Response *http.Response
	Err      error
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.Response, m.Err
}

func NewMockHTTPClient(body string, status int) *http.Client {
	return &http.Client{
		Transport: &MockRoundTripper{
			Response: &http.Response{
				StatusCode: status,
				Body:       io.NopCloser(bytes.NewBufferString(body)),
				Header:     make(http.Header),
			},
		},
	}
}

