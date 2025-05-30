package handler

import (
	"cli-crud/internal/auth"
	"cli-crud/internal/data/session"
	"cli-crud/pkg/io"
	"cli-crud/types"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func UpdateService(args []string, session *session.Session) {
	if !auth.CheckAuth(session) {
		return
	}

	headers := map[string]string {
		"Authorization": "Bearer " + session.AuthToken,
	}

	id, name, description, err := parseUpdateArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	service := types.Service {
		NAME:         name,
		DESCRIPTION:  description,
	}

	url := fmt.Sprintf("http://localhost:3000/service/%d", id)

	updatedService, err := io.DoJSONRequest[types.Service]("PUT", url, headers, service)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Updated successfully!")
	fmt.Printf("Id: %d\n", updatedService.ID)
}

func parseUpdateArgs(args []string) (int, string, string, error) {
	if len(args) < 3 {
		return 0, "", "", fmt.Errorf(`usage: update-service <id> "<name>" "<description>"`)
	}
	
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, "", "", fmt.Errorf("invalid ID: %v", err)
	}


	raw := strings.Join(args[1:], " ")

	re := regexp.MustCompile(`"([^"]+)"`)
	matches := re.FindAllStringSubmatch(raw, -1)

	if len(matches) < 2 {
		return 0, "", "", fmt.Errorf("please wrap both service name and description in quotes")

	}
		return id, matches[0][1], matches[1][1], nil
	}