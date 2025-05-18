package client

import (
	"errors"
)

type Client struct {
	BaseURL       string
	MarketplaceID string
}

func NewClient(countryCode string) (*Client, error) {
	reg, ok := countryRegionMap[countryCode]
	if !ok {
		return nil, ErrRegionNotFound
	}
	return &Client{
		BaseURL:       reg.BaseURL,
		MarketplaceID: reg.MarketplaceID,
	}, nil
}

var ErrRegionNotFound error = errors.New("geven country code was not found in regions")
