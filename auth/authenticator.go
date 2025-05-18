package auth

import (
    "context"
)

type Authenticator interface {
    GetAccessToken(ctx context.Context) (string, error)
}
