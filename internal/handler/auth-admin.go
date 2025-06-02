package handler

import (
	"fmt"
	"cli-crud/pkg/io"
	"cli-crud/internal/data/session"
	"cli-crud/types"
	"time"
)

type AuthResult struct {
	Response *types.LoginResponse
	Error    error
}

func AuthAdmin(args []string, sess *session.Session) {
	// Check if the user has provided the correct number of arguments
	if len(args) < 2 {
		fmt.Println("Usage: auth-admin <your email> <your password>")
		return
	}

	// Prepare the payload with user credentials
	payload := types.LoginRequest{
		Email:    args[0],
		Password: args[1],
	}

	// Start measuring the time for the operation
	startTime := time.Now()

	// Create a channel to receive the result from the goroutine
	resultChan := make(chan AuthResult, 1)

	// Launch the HTTP request in a goroutine for concurrent processing
	go func() {
		// Perform the HTTP request to authenticate the admin
		response, err := io.DoJSONRequest[types.LoginResponse](
			"POST",                  // HTTP method
			"http://localhost:3000/login", // Your backend URL
			nil,                     // Any custom headers (none in this case)
			payload,                 // The login payload containing the credentials
		)

		// Send the result or error back through the channel
		resultChan <- AuthResult{Response: response, Error: err}
	}()

	// Wait for the result from the goroutine
	result := <-resultChan

	if result.Error != nil {
		fmt.Println("Login failed:", result.Error)
		return
	}

	duration := time.Since(startTime)
	fmt.Printf("Auth operation completed in %v\n", duration)

	expiresAt, err := session.ExtractExpiryFromToken(result.Response.Token)
	if err == nil {
		sess.ExpiresAt = expiresAt
	}

	sess.AuthToken = result.Response.Token
	sess.UserEmail = payload.Email

	fmt.Println("Admin authenticated.")
}
// use sync.WaitGroup nxt time