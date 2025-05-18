package endpoint_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chiyonn/spapi/client"
	"github.com/chiyonn/spapi/endpoint"
)

// DummyRateLimitManager mocks Register behavior
type DummyRateLimitManager struct {
	RegisterFunc func(key string, rate float64, burst int) error
	WaitFunc     func(ctx context.Context, key string) error
}

func (d *DummyRateLimitManager) Register(key string, rate float64, burst int) error {
	if d.RegisterFunc != nil {
		return d.RegisterFunc(key, rate, burst)
	}
	return nil
}

func (d *DummyRateLimitManager) Wait(ctx context.Context, key string) error {
	if d.WaitFunc != nil {
		return d.WaitFunc(ctx, key)
	}
	return nil // default no-op for testing
}

// Dummy HTTP client to simulate API server
func dummyServer(t *testing.T, statusCode int, responseBody string) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		_, err := w.Write([]byte(responseBody))
		if err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	})
	return httptest.NewServer(handler)
}

func TestDo_Success(t *testing.T) {
	server := dummyServer(t, 200, `{"status":"ok"}`)
	defer server.Close()

	cli := &client.Client{
		HttpClient:        http.DefaultClient,
		BaseURL:           server.URL,
		RateLimitManager:  &DummyRateLimitManager{},
	}

	endpoint := &endpoint.APIEndpoint{
		Client:  cli,
		Method:  http.MethodGet,
		Path:    "/test",
		BuildReq: func() (*http.Request, error) {
			return http.NewRequest(http.MethodGet, cli.BaseURL+"/test", nil)
		},
		ParseResp: func(resp *http.Response) (any, error) {
			body, _ := io.ReadAll(resp.Body)
			return string(body), nil
		},
	}

	result, err := endpoint.Do(context.Background())
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	expected := `{"status":"ok"}`
	if result != expected {
		t.Errorf("Do() got = %v, want = %v", result, expected)
	}
}

func TestDo_BuildReqError(t *testing.T) {
	mockClient := &client.Client{
		HttpClient: http.DefaultClient,
		BaseURL:    "https://example.com",
		RateLimitManager: &DummyRateLimitManager{
			WaitFunc: func(ctx context.Context, key string) error {
				return nil // noop
			},
		},
	}

	endpoint := &endpoint.APIEndpoint{
		Client: mockClient,
		BuildReq: func() (*http.Request, error) {
			return nil, errors.New("build error")
		},
		ParseResp: func(resp *http.Response) (any, error) {
			return nil, nil
		},
		RateKey: "test.endpoint",
	}

	_, err := endpoint.Do(context.Background())
	if err == nil || err.Error() != "build error" {
		t.Errorf("Expected build error, got %v", err)
	}
}

func TestNewEndpoint_Success(t *testing.T) {
	cli := &client.Client{
		RateLimitManager: &DummyRateLimitManager{
			RegisterFunc: func(key string, rate float64, burst int) error {
				if key != "test-key" || rate != 1.0 || burst != 2 {
					return errors.New("invalid registration")
				}
				return nil
			},
		},
	}

	ep, err := endpoint.NewEndpoint(cli, "GET", "/path", 1.0, 2, "test-key")
	if err != nil {
		t.Fatalf("NewEndpoint() unexpected error: %v", err)
	}

	if ep.Path != "/path" || ep.Method != "GET" {
		t.Errorf("NewEndpoint() incorrect values: %+v", ep)
	}
}

func TestNewEndpoint_RegisterFails(t *testing.T) {
	cli := &client.Client{
		RateLimitManager: &DummyRateLimitManager{
			RegisterFunc: func(string, float64, int) error {
				return errors.New("fail")
			},
		},
	}

	_, err := endpoint.NewEndpoint(cli, "GET", "/path", 1.0, 2, "key")
	if err != endpoint.ErrInitRateLimitManager {
		t.Errorf("Expected ErrInitRateLimitManager, got %v", err)
	}
}
