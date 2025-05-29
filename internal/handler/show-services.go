package handler

import (
	"cli-crud/internal/data/session"
	"cli-crud/pkg/io"
	"fmt"
	"time"
)

type Service struct {
	ID int `json:"id"`
	NAME string `json:"name"`
	DESCRIPTION string `json:"description"`
}

func GetServiceList(args []string, session *session.Session) {
	if session.AuthToken == "" || time.Now().After(session.ExpiresAt) {
		fmt.Println("Session expired. Please log in")
		return
	}
	headers := map[string]string {
		"Authorization": "Bearer " + session.AuthToken,
	}

	services, err := io.DoJSONRequest[[]Service]("GET", "http://localhost:3000/services", headers, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	
	fmt.Println("Services:")
	for _, service := range *services {
		fmt.Printf("- %d: %s (%s)\n", service.ID, service.NAME, service.DESCRIPTION)
	}
}