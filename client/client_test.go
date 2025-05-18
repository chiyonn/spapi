package client_test

import (
	"testing"

	"github.com/chiyonn/spapi/client"
)

func TestNewClient_Success(t *testing.T) {
	rlm := client.NewRateLimitManager()
	c, err := client.NewClient("JP", rlm)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if c.BaseURL != "https://sellingpartnerapi-fe.amazon.com" {
		t.Errorf("Unexpected BaseURL: got %s", c.BaseURL)
	}

	if c.MarketplaceID != "A1VC38T7YXB528" {
		t.Errorf("Unexpected MarketplaceID: got %s", c.MarketplaceID)
	}
}

func TestNewClient_Failure(t *testing.T) {
	rlm := client.NewRateLimitManager()
	c, err := client.NewClient("XX", rlm)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err != client.ErrRegionNotFound {
		t.Errorf("Expected ErrRegionNotFound, got %v", err)
	}

	if c != nil {
		t.Errorf("Expected client to be nil on error, got %+v", c)
	}
}
