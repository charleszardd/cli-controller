package auth

import (
	"cli-crud/internal/data/session"
	"fmt"
	"time"
)

func CheckAuth(s *session.Session) bool {
	if s.AuthToken == "" || time.Now().After(s.ExpiresAt) {
		fmt.Println("Session expired. Please log in")
		return false
	}

	return true

}