package endpoint_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/chiyonn/spapi/client"
	"github.com/chiyonn/spapi/endpoint"
)

type mockRateLimiter struct {
	waitCalled     bool
	registerCalled bool
	failWait       bool
	failRegister   bool
}

func (m *mockRateLimiter) Wait(ctx context.Context, key string) error {
	m.waitCalled = true
	if m.failWait {
		return errors.New("wait failed")
	}
	return nil
}

func (m *mockRateLimiter) Register(key string, rate float64, burst int) error {
	m.registerCalled = true
	if m.failRegister {
		return errors.New("register failed")
	}
	return nil
}

type mockAuth struct {
	token string
	err   error
}

func (m *mockAuth) GetAccessToken(ctx context.Context) (string, error) {
	return m.token, m.err
}

type mockRoundTripper struct {
	resp *http.Response
	err  error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.resp, m.err
}

func newMockClient(token string, body string, statusCode int) *client.Client {
	return &client.Client{
		RateLimitManager: &mockRateLimiter{},
		Auth:             &mockAuth{token: token},
		HTTPClient: &http.Client{
			Transport: &mockRoundTripper{
				resp: &http.Response{
					StatusCode: statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(body)),
				},
			},
		},
	}
}

func TestNewEndpointValidationErrors(t *testing.T) {
	c := newMockClient("token", "{}", 200)

	_, err := endpoint.NewEndpoint(nil, "GET", "/path", 1.0, 10, "key")
	if err != endpoint.ErrEmptyClientDetected {
		t.Errorf("expected ErrEmptyClientDetected, got %v", err)
	}

	_, err = endpoint.NewEndpoint(c, "", "/path", 1.0, 10, "key")
	if err != endpoint.ErrEmptyMethodDetected {
		t.Errorf("expected ErrEmptyMethodDetected, got %v", err)
	}

	_, err = endpoint.NewEndpoint(c, "GET", "", 1.0, 10, "key")
	if err != endpoint.ErrEmptyPathDetected {
		t.Errorf("expected ErrEmptyPathDetected, got %v", err)
	}

	_, err = endpoint.NewEndpoint(c, "GET", "/path", 1.0, 10, "")
	if err != endpoint.ErrEmptyRateKeyDetected {
		t.Errorf("expected ErrEmptyRateKeyDetected, got %v", err)
	}

	_, err = endpoint.NewEndpoint(c, "GET", "/path", 0.0, 10, "key")
	if err != endpoint.ErrEmptyRateDetected {
		t.Errorf("expected ErrEmptyRateDetected, got %v", err)
	}

	_, err = endpoint.NewEndpoint(c, "GET", "/path", 1.0, 0, "key")
	if err != endpoint.ErrEmptyBurstRateDetected {
		t.Errorf("expected ErrEmptyBurstRateDetected, got %v", err)
	}
}

func TestDoSuccess(t *testing.T) {
	c := newMockClient("token123", `{"success":true}`, 200)

	ep, err := endpoint.NewEndpoint(c, "GET", "/test", 1.0, 10, "rate-key")
	if err != nil {
		t.Fatalf("unexpected error from NewEndpoint: %v", err)
	}

	ep.BuildReq = func() (*http.Request, error) {
		req, _ := http.NewRequest("GET", "https://example.com", nil)
		return req, nil
	}

	ep.ParseResp = func(resp *http.Response) (any, error) {
		body, _ := io.ReadAll(resp.Body)
		return string(body), nil
	}

	result, err := ep.Do(context.Background())
	if err != nil {
		t.Fatalf("Do returned unexpected error: %v", err)
	}

	if result != `{"success":true}` {
		t.Errorf("unexpected response: %v", result)
	}
}

func TestDoFailures(t *testing.T) {
	c := newMockClient("token123", "{}", 200)
	rl := &mockRateLimiter{failWait: true}
	c.RateLimitManager = rl

	ep, _ := endpoint.NewEndpoint(c, "GET", "/fail", 1.0, 10, "rk")
	ep.BuildReq = func() (*http.Request, error) {
		req, _ := http.NewRequest("GET", "https://example.com", nil)
		return req, nil
	}
	ep.ParseResp = func(resp *http.Response) (any, error) {
		return nil, nil
	}

	// Rate limiter fails
	_, err := ep.Do(context.Background())
	if err == nil || err.Error() != "wait failed" {
		t.Errorf("expected wait failure, got %v", err)
	}
}
