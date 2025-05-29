package handler

import (
	"fmt"
	"cli-crud/pkg/io"
	"cli-crud/internal/data/session"
)

type LoginRequest struct  {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// serve as a controller for the auth-admin command
func AuthAdmin(args []string, sess *session.Session) {
	if len(args) < 2 {
		fmt.Println("Usage: auth-admin <your email> <your password>")
		return
	}

	payload := LoginRequest{
		Email: args[0],
		Password: args[1],
	}

	response, err := io.DoJSONRequest[LoginResponse](
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