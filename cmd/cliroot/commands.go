package main

import (
	"cli-crud/internal/handler"
	"cli-crud/internal/data/session"
)

const (
	CmdHelp = "help"
	CmdExit = "EXIT"
)

type CmdConfig struct {
	minParams  int
	params     string
	handler    func([]string,*session.Session)

}

var Config = map[string]CmdConfig{
	// CmdHelp:              {0, "", nil},

	"AUTH-ADMIN":         {2, Auth, handler.AuthAdmin},
	// "post-list":          {1, JsonPath, handler.PostList},
	"SHOW-USERS":           {0, "Display all the users from the database. You need to log in first, though", handler.GetUsersList},
	"SHOW-SERVICES":        {0, "Display all the services from the database. You need to log in first, though", handler.GetServiceList},
	// "update-list-id":     {2, ListID + " " + JsonPath, handler.UpdateList},
	// "delete-list":        {1, ListID, handler.DeleteList},
	// "delete-list all":    {0, "", handler.DeleteListAll},
	// "show-users":         {},
	CmdExit:              {0, "Exit program", nil},
}