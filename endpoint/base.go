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
	endpoint := APIEndpoint{
		Client: client,
		Method: method,
		Path:   path,
	}
	err := endpoint.Client.RateLimitManager.Register(key, rate, burst)
	if err != nil {
		return nil, ErrInitRateLimitManager
	}
	return &endpoint, nil
}

var ErrInitRateLimitManager = errors.New("failed to initialize RateLimitManager")
