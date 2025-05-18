package auth

// AuthConfig holds the credentials required for OAuth2 token generation.
type AuthConfig struct {
	RefreshToken string
	ClientID     string
	ClientSecret string
}

// NewAuthConfig validates and returns a new AuthConfig.
// Returns ErrConfigNotSet if any field is empty.
func NewAuthConfig(rt string, cid string, cs string) (*AuthConfig, error) {
	cfg := &AuthConfig{
		RefreshToken: rt,
		ClientID:     cid,
		ClientSecret: cs,
	}
	if cfg.RefreshToken == "" || cfg.ClientID == "" || cfg.ClientSecret == "" {
		return nil, ErrConfigNotSet
	}
	return cfg, nil
}
