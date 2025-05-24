package testutil

import (
	"context"
	"testing"

	"github.com/chiyonn/spapi/client"
)

type noopAuthenticator struct{}

func (n *noopAuthenticator) GetAccessToken(ctx context.Context) (string, error) {
	return "dummy-access-token", nil
}

func NewMockedClient(t *testing.T, handler RoundTripFunc) *client.Client {
	t.Helper()

	return &client.Client{
		BaseURL:          "https://mock.api",
		HTTPClient:       NewMockHTTPClient(handler),
		RateLimitManager: client.NewRateLimitManager(),
		Auth:             &noopAuthenticator{},
	}
}

