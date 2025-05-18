package client

import (
	"net/http"

	"github.com/chiyonn/spapi/auth"
)

type Client struct {
    HttpClient       *http.Client
    BaseURL          string
    MarketplaceID    string
    RateLimitManager RateLimitManager
    Auth             auth.Authenticator
}

func NewClient(cc string, cfg *auth.AuthConfig, rlm RateLimitManager) (*Client, error) {
    reg, ok := countryRegionMap[cc]
    if !ok {
        return nil, ErrRegionNotFound
    }

    httpClient := &http.Client{} // or injected externally
    authenticator := auth.NewOAuth2Authenticator(httpClient, cfg)

    return &Client{
        HttpClient:       httpClient,
        BaseURL:          reg.BaseURL,
        MarketplaceID:    reg.MarketplaceID,
        RateLimitManager: rlm,
        Auth:             authenticator,
    }, nil
}
