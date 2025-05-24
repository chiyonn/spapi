package client

import (
	"net/http"

	"github.com/chiyonn/spapi/auth"
)

type Client struct {
    HTTPClient       *http.Client
    BaseURL          string
    MarketplaceID    string
    RateLimitManager RateLimitManager
    Auth             auth.Authenticator
}

func NewClient(cli *http.Client, cc string, cfg *auth.AuthConfig, rlm RateLimitManager) (*Client, error) {
    reg, ok := countryRegionMap[cc]
    if !ok {
        return nil, ErrRegionNotFound
    }

    authenticator := auth.NewOAuth2Authenticator(cli, cfg)

    return &Client{
        HTTPClient:       cli,
        BaseURL:          reg.BaseURL,
        MarketplaceID:    reg.MarketplaceID,
        RateLimitManager: rlm,
        Auth:             authenticator,
    }, nil
}
