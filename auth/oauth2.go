package auth

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "strings"
    "sync"
    "time"
)

const tokenEndpoint = "https://api.amazon.com/auth/o2/token"

type OAuth2Authenticator struct {
    Client       *http.Client
    RefreshToken string
    ClientID     string
    ClientSecret string

    mu          sync.Mutex
    accessToken string
    expiresAt   time.Time
}

func NewOAuth2Authenticator(client *http.Client, cfg *AuthConfig) *OAuth2Authenticator {
	return &OAuth2Authenticator{
		Client:       client,
		RefreshToken: cfg.RefreshToken,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
	}
}

func (a *OAuth2Authenticator) GetAccessToken(ctx context.Context) (string, error) {
    a.mu.Lock()
    defer a.mu.Unlock()

    if time.Until(a.expiresAt) > 2*time.Minute {
        return a.accessToken, nil // still valid
    }

    form := url.Values{}
    form.Set("grant_type", "refresh_token")
    form.Set("refresh_token", a.RefreshToken)
    form.Set("client_id", a.ClientID)
    form.Set("client_secret", a.ClientSecret)

    req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenEndpoint, strings.NewReader(form.Encode()))
    if err != nil {
        return "", fmt.Errorf("failed to create token request: %w", err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    resp, err := a.Client.Do(req)
    if err != nil {
        return "", fmt.Errorf("token request failed: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return "", fmt.Errorf("token fetch failed: %d %s", resp.StatusCode, string(body))
    }

	var res struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	token := NewToken(res.AccessToken, res.ExpiresIn)

	a.accessToken = token.AccessToken
	a.expiresAt = token.ExpiresAt

	return token.AccessToken, nil
}

