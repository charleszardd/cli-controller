package handler

import (
	"cli-crud/internal/data/session"
	"cli-crud/internal/auth"
	"cli-crud/pkg/io"
	"cli-crud/types"
	"fmt"
	"strings"
)

func GetServiceList(args []string, session *session.Session) {
	
	if !auth.CheckAuth(session) {
		return
	}

	headers := map[string]string {
		"Authorization": "Bearer " + session.AuthToken,
	}

	services, err := io.DoJSONRequest[[]types.Service]("GET", "http://localhost:3000/services", headers, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Services:")
	for _, service := range *services {
		name := strings.Trim(service.NAME, "\"")
		description := strings.Trim(service.DESCRIPTION, "\"")
		fmt.Printf("- %d: %s (%s)\n", service.ID, name, description)
	}
}