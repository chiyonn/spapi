package client

import (
	"errors"
)

var (
	ErrRegionNotFound     error = errors.New("given country code was not found in regions")
)

