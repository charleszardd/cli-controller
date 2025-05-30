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

	"AUTH-ADMIN"          :    {2, Auth, handler.AuthAdmin},
	"SHOW-USERS"          :    {0, DisplayAllUsersMessage, handler.GetUsersList},
	"SHOW-SERVICES"       :    {0, DisplayAllServicesMessage, handler.GetServiceList},
	"CREATE-SERVICE"      :    {2, ServiceFields, handler.PostService},
	"DELETE-ALL-SERVICES" :    {0, ShowDeleteAllServicesMessage, handler.DeleteAllServices},
	"DELETE-SERVICE-ID"   :    {1, ShowDeleteServiceIDMessage, handler.DeleteById},
	"UPDATE-SERVICE"      :    {2, ShowUpdateServiceMessage, handler.UpdateService},
	CmdExit               :    {0, "Exit program", nil},
}