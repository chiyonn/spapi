package endpoint

import (
	"context"
	"errors"
	"net/http"

	"github.com/chiyonn/spapi/client"
)

type APIEndpoint struct {
	c    	  *client.Client
	Path      string
	Method    string
	Rate      float64
	Burst     int
	RateKey   string
	BuildReq  func() (*http.Request, error)
	ParseResp func(*http.Response) (any, error)
}

func (ep *APIEndpoint) Do(ctx context.Context) (any, error) {
	if ep.c == nil || ep.c.HttpClient == nil {
		return nil, errors.New("client or HttpClient is nil")
	}

	if err := ep.c.RateLimitManager.Wait(ctx, ep.RateKey); err != nil {
		return nil, err
	}

	req, err := ep.BuildReq()
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, errors.New("BuildReq returned nil request")
	}

	token, err := ep.c.Auth.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-amz-access-token", token)

	resp, err := ep.c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ep.ParseResp(resp)
}

func NewEndpoint(client *client.Client, method string, path string, rate float64, burst int, key string) (*APIEndpoint, error) {
	if client == nil {
		return nil, ErrEmptyClientDetected
	}
	if method == "" {
		return nil, ErrEmptyMethodDetected
	}
	if path == "" {
		return nil, ErrEmptyPathDetected
	}
	if key == "" {
		return nil, ErrEmptyRateKeyDetected
	}
	if rate <= 0 {
		return nil, ErrEmptyRateDetected
	}
	if burst <= 0 {
		return nil, ErrEmptyBurstRateDetected
	}

	endpoint := APIEndpoint{
		c:  client,
		Method:  method,
		Path:    path,
		RateKey: key,
	}

	err := endpoint.c.RateLimitManager.Register(key, rate, burst)
	if err != nil {
		return nil, ErrInitRateLimitManager
	}
	return &endpoint, nil
}

var (
	ErrInitRateLimitManager   = errors.New("failed to initialize RateLimitManager")
	ErrEmptyClientDetected    = errors.New("client must not be nil")
	ErrEmptyMethodDetected    = errors.New("method must not be nil")
	ErrEmptyPathDetected      = errors.New("path must not be nil")
	ErrEmptyRateKeyDetected   = errors.New("rate key must not be nil")
	ErrEmptyRateDetected      = errors.New("rate must not be nil")
	ErrEmptyBurstRateDetected = errors.New("rate burst not be nil")
)
