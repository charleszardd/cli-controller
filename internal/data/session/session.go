package session

import (
	"time"
)

type Session struct {
	AuthToken string
	UserEmail string
	ExpiresAt time.Time
}