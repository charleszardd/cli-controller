package handler

import (
	"cli-crud/internal/auth"
	"cli-crud/internal/data/session"
	"cli-crud/pkg/io"
	"cli-crud/types"
	"fmt"
)

func DeleteAllServices(args []string, session *session.Session) {
	if !auth.CheckAuth(session) {
		return
	}

	headers := map[string]string{
		"Authorization": "Bearer " + session.AuthToken,
	}

	response, err := io.DoJSONRequest[types.ServiceMessageResponse]("DELETE", "http://localhost:3000/services", headers, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(response.Message)

}