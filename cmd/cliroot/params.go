package main

import (
	"fmt"
)

const  (
	ListID 							=    "<list id>"
	ShowCreateServiceMessage 		=    `create-service "<service name>" "<service description>"`
	DisplayAllServicesMessage 		=    "Display all the services from the database"
	DisplayAllUsersMessage 			=    "Display all the users from the database"
	Auth 							=    "auth-admin <email> <password>"
	ShowUpdateServiceMessage 		=    "Update a specific service from the database <choose id> <new service name> <new service description>"
	ShowDeleteAllServicesMessage 	=    "Delete all the services from the database"
	ShowDeleteServiceIDMessage 		=    "Delete a specific service from the database"
)


func validateCommandParams(conf CmdConfig, params []string) error {
	if len(params) < conf.minParams {
		return fmt.Errorf("usage: %s",  conf.params)
	}
	return nil
}
