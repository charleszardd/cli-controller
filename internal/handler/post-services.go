package handler

import (
	"cli-crud/internal/data/session"
	"cli-crud/internal/auth"
	"cli-crud/pkg/io"
	"cli-crud/types"
	"fmt"
	"regexp"
	"strings"
)

func PostService(args []string, session *session.Session) {
	if !auth.CheckAuth(session) {
		return
	}

	name, description, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	headers := map[string]string{
		"Authorization": "Bearer " + session.AuthToken,
	}

	service := types.Service{
		NAME:        name,
		DESCRIPTION: description,
	}

	createdService, err := io.DoJSONRequest[types.Service]("POST", "http://localhost:3000/service", headers, service)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Successfully created service:")
	fmt.Printf("Id: %d\n", createdService.ID)
}

func parseArgs(args []string) (string, string, error) {
	if len(args) < 2 {
		return "", "", fmt.Errorf("usage: create-services %s %s", "<service name>", "<service description>")
	}

	raw := strings.Join(args, " ")

	re := regexp.MustCompile(`"([^"]+)"`)
	matches := re.FindAllStringSubmatch(raw, -1)

	if len(matches) < 2 {
		return "", "", fmt.Errorf("please wrap both service name and description in quotes")
	}

	return matches[0][1], matches[1][1], nil
}

