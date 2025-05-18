package client

import (
	"errors"
	"net/http"
)

type Client struct {
	HttpClient 		 *http.Client
	BaseURL          string
	MarketplaceID    string
	RateLimitManager RateLimitManager
}

func NewClient(cc string, rlm RateLimitManager) (*Client, error) {
	reg, ok := countryRegionMap[cc]
	if !ok {
		return nil, ErrRegionNotFound
	}
	return &Client{
		HttpClient: &http.Client{},
		BaseURL:          reg.BaseURL,
		MarketplaceID:    reg.MarketplaceID,
		RateLimitManager: rlm,
	}, nil
}

var ErrRegionNotFound error = errors.New("geven country code was not found in regions")
