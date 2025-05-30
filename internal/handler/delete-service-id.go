package handler

import (
	"cli-crud/internal/auth"
	"cli-crud/internal/data/session"
	"cli-crud/types"
	"cli-crud/pkg/io"
	"fmt"
	"strconv"
)

func DeleteById(args []string, session *session.Session) {
	if !auth.CheckAuth(session) {
		return
	}

	headers := map[string]string{
		"Authorization": "Bearer " + session.AuthToken,
	}

	id, err := parseDeleteArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	url := fmt.Sprintf("http://localhost:3000/service/%d", id)

	deletedService, err := io.DoJSONRequest[types.Service]("DELETE", url, headers, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Deleted successfully!")
	fmt.Printf("Id: %d deleted\n", deletedService.ID)
}

func  parseDeleteArgs(args []string) (int, error) {
	if len(args) < 1 {
		return 0, fmt.Errorf(`usage: delete-service <service id>`)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, fmt.Errorf("invalid ID: %v", err)
	}
	return id, nil
}