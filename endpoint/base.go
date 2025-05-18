package endpoint

import (
	"context"
	"errors"
	"net/http"

	"github.com/chiyonn/spapi/client"
)

type APIEndpoint struct {
	Path      string
	Method    string
	Rate      float64
	Burst     int
	RateKey   string
	Client    *client.Client
	BuildReq  func() (*http.Request, error)
	ParseResp func(*http.Response) (any, error)
}

func (ep *APIEndpoint) Do(ctx context.Context) (any, error) {
	if err := ep.Client.RateLimitManager.Wait(ctx, ep.RateKey); err != nil {
		return nil, err
	}

	req, err := ep.BuildReq()
	if err != nil {
		return nil, err
	}

	resp, err := ep.Client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ep.ParseResp(resp)
}

func NewEndpoint(client *client.Client, method string, path string, rate float64, burst int, key string) (*APIEndpoint, error) {
	if client == nil {
		return nil, errors.New("client must not be nil")
	}
	if method == "" {
		return nil, errors.New("method must not be empty")
	}
	if path == "" {
		return nil, errors.New("path must not be empty")
	}
	if key == "" {
		return nil, errors.New("rate key must not be empty")
	}
	if rate <= 0 {
		return nil, errors.New("rate must be greater than 0")
	}
	if burst <= 0 {
		return nil, errors.New("burst must be greater than 0")
	}

	endpoint := APIEndpoint{
		Client:  client,
		Method:  method,
		Path:    path,
		RateKey: key,
	}

	err := endpoint.Client.RateLimitManager.Register(key, rate, burst)
	if err != nil {
		return nil, ErrInitRateLimitManager
	}
	return &endpoint, nil
}

var ErrInitRateLimitManager = errors.New("failed to initialize RateLimitManager")
