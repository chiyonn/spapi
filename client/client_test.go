package client_test

import (
	"testing"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/chiyonn/spapi/client"
)

func TestNewClient_Success(t *testing.T) {
	cfg, err := client.NewClientConfig("refresh_token", "client_id", "client_secret")
	rlm := client.NewRateLimitManager()
	c, err := client.NewClient("JP", cfg, rlm)

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
	cfg, err := client.NewClientConfig("refresh_token", "client_id", "client_secret")
	rlm := client.NewRateLimitManager()
	c, err := client.NewClient("XX", cfg, rlm)
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

func TestGetAccessToken_Success(t *testing.T) {
	// モックレスポンス
	tokenResponse := map[string]interface{}{
		"access_token": "mock_token",
		"expires_in":   3600,
	}
	respBody, _ := json.Marshal(tokenResponse)

	// モックサーバー
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if err := r.ParseForm(); err != nil {
			t.Errorf("Failed to parse form: %v", err)
		}
		if r.FormValue("grant_type") != "refresh_token" {
			t.Errorf("Unexpected grant_type: %s", r.FormValue("grant_type"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer server.Close()

	// トークンエンドポイント上書き
	restore := client.OverrideTokenEndpoint(server.URL)
	defer restore()

	// クライアント構築
	cfg, _ := client.NewClientConfig("refresh_token", "client_id", "client_secret")
	rlm := client.NewRateLimitManager()
	c, err := client.NewClient("JP", cfg, rlm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	token, err := c.GetAccessToken(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == nil || *token != "mock_token" {
		t.Errorf("Expected 'mock_token', got %v", token)
	}
}

func TestGetAccessToken_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}))
	defer server.Close()

	restore := client.OverrideTokenEndpoint(server.URL)
	defer restore()

	cfg, _ := client.NewClientConfig("refresh_token", "client_id", "client_secret")
	c, _ := client.NewClient("JP", cfg, client.NewRateLimitManager())

	token, err := c.GetAccessToken(context.Background())
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if token != nil {
		t.Errorf("Expected nil token, got %v", token)
	}
}

func TestGetAccessToken_AlreadyValid(t *testing.T) {
	c := &client.Client{
		HttpClient:        http.DefaultClient,
		BaseURL:           "https://dummy",
		RateLimitManager:  client.NewRateLimitManager(),
		refreshToken:      "dummy_refresh",
		clientID:          "dummy_id",
		clientSecret:      "dummy_secret",
		accessToken:       "cached_token",
		expiresAt:         time.Now().Add(10 * time.Minute),
	}

	token, err := c.GetAccessToken(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == nil || *token != "cached_token" {
		t.Errorf("Expected 'cached_token', got %v", token)
	}
}

func TestGetAccessToken_DecodeError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "not-json")
	}))
	defer server.Close()

	restore := client.OverrideTokenEndpoint(server.URL)
	defer restore()

	cfg, _ := client.NewClientConfig("refresh_token", "client_id", "client_secret")
	c, _ := client.NewClient("JP", cfg, client.NewRateLimitManager())

	token, err := c.GetAccessToken(context.Background())
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if token != nil {
		t.Errorf("Expected nil token, got %v", token)
	}
}
