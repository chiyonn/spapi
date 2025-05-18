package auth

import (
	"errors"
)

var (
	ErrAccessTokenRequest error = errors.New("failed to create token request")
	ErrConfigNotSet error = errors.New("clint config must set all")
)
