package auth

import "time"

// Token represents an access token with its expiration.
type Token struct {
    AccessToken string    `json:"access_token"`
    ExpiresAt   time.Time `json:"expires_at"`
}

// NewToken creates a Token from an access_token and expires_in (in seconds).
func NewToken(accessToken string, expiresIn int) *Token {
    expiresAt := time.Now().Add(time.Duration(expiresIn-60) * time.Second) // 60s buffer
    return &Token{
        AccessToken: accessToken,
        ExpiresAt:   expiresAt,
    }
}

// IsValid returns true if the token is still valid with an optional buffer.
func (t *Token) IsValid(buffer time.Duration) bool {
    return time.Until(t.ExpiresAt) > buffer
}

