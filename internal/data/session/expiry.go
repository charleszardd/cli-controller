package session

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

func ExtractExpiryFromToken(token string) (time.Time, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return time.Time{}, errors.New("invalid token format")
	}

	payloadPart := parts[1]

	switch len(payloadPart) % 4 {
	case 2:
		payloadPart += "=="
	case 3:
		payloadPart += "="
		
	}

	payloadBytes, err :=base64.URLEncoding.DecodeString(payloadPart)
	if err != nil {
		return time.Time{}, err
	}

	var claims struct {
		Exp int64 `json:"exp"`
	}

	if  err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return time.Time{}, err
	}

	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return time.Time{}, err
	}

	return time.Unix(claims.Exp, 0), nil
}