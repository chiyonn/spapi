package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetAccessToken_Success(t *testing.T) {
	mockToken := "test_token"
	mockExpiresIn := 3600

	// Set up a fake token server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Fatalf("failed to parse form: %v", err)
		}
		if r.FormValue("grant_type") != "refresh_token" {
			t.Errorf("unexpected grant_type: %s", r.FormValue("grant_type"))
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"access_token": "%s", "expires_in": %d}`, mockToken, mockExpiresIn)
	}))
	defer server.Close()

	// Replace endpoint with mock server
	originalEndpoint := tokenEndpoint
	tokenEndpoint = server.URL
	defer func() { tokenEndpoint = originalEndpoint }()

	auth := &OAuth2Authenticator{
		Client:       server.Client(),
		RefreshToken: "rt",
		ClientID:     "cid",
		ClientSecret: "cs",
	}

	token, err := auth.GetAccessToken(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token != mockToken {
		t.Errorf("expected token %s, got %s", mockToken, token)
	}
}

func TestGetAccessToken_CachesToken(t *testing.T) {
	auth := &OAuth2Authenticator{
		Client:       nil, // should not be used
		RefreshToken: "rt",
		ClientID:     "cid",
		ClientSecret: "cs",
		accessToken:  "cached_token",
		expiresAt:    time.Now().Add(10 * time.Minute),
	}

	token, err := auth.GetAccessToken(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token != "cached_token" {
		t.Errorf("expected cached token, got %s", token)
	}
}

func TestGetAccessToken_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad request", http.StatusBadRequest)
	}))
	defer server.Close()

	tokenEndpoint = server.URL
	defer func() { tokenEndpoint = "https://api.amazon.com/auth/o2/token" }()

	auth := &OAuth2Authenticator{
		Client:       server.Client(),
		RefreshToken: "rt",
		ClientID:     "cid",
		ClientSecret: "cs",
	}

	_, err := auth.GetAccessToken(context.Background())
	if err == nil || !strings.Contains(err.Error(), "token fetch failed") {
		t.Errorf("expected token fetch failure, got: %v", err)
	}
}

func TestGetAccessToken_DecodeError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{invalid json}`)
	}))
	defer server.Close()

	tokenEndpoint = server.URL
	defer func() { tokenEndpoint = "https://api.amazon.com/auth/o2/token" }()

	auth := &OAuth2Authenticator{
		Client:       server.Client(),
		RefreshToken: "rt",
		ClientID:     "cid",
		ClientSecret: "cs",
	}

	_, err := auth.GetAccessToken(context.Background())
	if err == nil || !strings.Contains(err.Error(), "failed to decode token response") {
		t.Errorf("expected decode error, got: %v", err)
	}
}

