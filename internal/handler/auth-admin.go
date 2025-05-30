package handler

import (
	"fmt"
	"cli-crud/pkg/io"
	"cli-crud/internal/data/session"
	"cli-crud/types"
)

// serve as a controller for the auth-admin command
func AuthAdmin(args []string, sess *session.Session) {
	if len(args) < 2 {
		fmt.Println("Usage: auth-admin <your email> <your password>")
		return
	}

	payload := types.LoginRequest{
		Email: args[0],
		Password: args[1],
	}

	response, err := io.DoJSONRequest[types.LoginResponse](
		"POST",
		"http://localhost:3000/login",
		nil,
		payload,
	)

	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}

	expiresAt, err := session.ExtractExpiryFromToken(response.Token)
	if err == nil {
		sess.ExpiresAt = expiresAt
	}

	sess.AuthToken = response.Token
	sess.UserEmail = payload.Email

	fmt.Println("Admin authenticated.")
}