package helpers

import "time"

const (
	MaxImageSize  = 1 << 20 // 1 MB
	SecretKey     = "secret"
	TokenDuration = 15 * time.Minute
)
