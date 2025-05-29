package handler

import (
	"cli-crud/pkg/io"
	"fmt"
	"time"

	"cli-crud/internal/data/session"
)

type User struct {
	ID int `json:"id"`
	NAME string `json:"name"`
	EMAIL string `json:"email"`
}

func GetUsersList(args []string, session *session.Session) {
	if session.AuthToken == "" || time.Now().After(session.ExpiresAt) {
		fmt.Println("Session expired. Please log in")
		return
	}
	headers := map[string]string {
		"Authorization": "Bearer " + session.AuthToken,
	}

	users, err := io.DoJSONRequest[[]User]("GET", "http://localhost:3000/users", headers, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Users:")
	for _, user := range *users {
		fmt.Printf("- %d: %s (%s)\n", user.ID, user.NAME, user.EMAIL)
	}
}