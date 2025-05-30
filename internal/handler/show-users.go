package handler

import (
	
	"fmt"

	"cli-crud/internal/data/session"
	"cli-crud/internal/auth"
	"cli-crud/types"
	"cli-crud/pkg/io"
)

func GetUsersList(args []string, session *session.Session) {
	
	if !auth.CheckAuth(session) {
		return
	}

	headers := map[string]string {
		"Authorization": "Bearer " + session.AuthToken,
	}

	users, err := io.DoJSONRequest[[]types.User]("GET", "http://localhost:3000/users", headers, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Users:")
	for _, user := range *users {
		fmt.Printf("- %d: %s (%s)\n", user.ID, user.NAME, user.EMAIL)
	}
}