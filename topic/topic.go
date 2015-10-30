package topic

import (
	"encoding/json"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var recursive bool

func init() {
	Cmd.AddCommand(cmdTopicList)
	Cmd.AddCommand(cmdTopicCreate)
	Cmd.AddCommand(cmdTopicDelete)
	Cmd.AddCommand(cmdTopicAddRoUser)
	Cmd.AddCommand(cmdTopicAddRwUser)
	Cmd.AddCommand(cmdTopicAddAdminUser)
	Cmd.AddCommand(cmdTopicDeleteRoUser)
	Cmd.AddCommand(cmdTopicDeleteRwUser)
	Cmd.AddCommand(cmdTopicDeleteAdminUser)
	Cmd.AddCommand(cmdTopicAddRoGroup)
	Cmd.AddCommand(cmdTopicAddRwGroup)
	Cmd.AddCommand(cmdTopicAddAdminGroup)
	Cmd.AddCommand(cmdTopicDeleteRoGroup)
	Cmd.AddCommand(cmdTopicDeleteRwGroup)
	Cmd.AddCommand(cmdTopicDeleteAdminGroup)
	Cmd.AddCommand(cmdTopicAddParameter)
	Cmd.AddCommand(cmdTopicDeleteParameter)
	Cmd.AddCommand(cmdTopicParameter)
}

// Cmd topic
var Cmd = &cobra.Command{
	Use:     "topic",
	Short:   "Topic commands: tatcli topic --help",
	Long:    "Topic commands: tatcli topic [command]",
	Aliases: []string{"t"},
}

type topicUsernameJSON struct {
	Topic     string `json:"topic"`
	Username  string `json:"username"`
	Recursive bool   `json:"recursive"`
}

func topicAddDeleteUsers(method string, topic string, users []string, path string) {
	for _, username := range users {
		t := topicUsernameJSON{topic, username, recursive}
		jsonStr, err := json.Marshal(t)
		internal.Check(err)
		if method == "PUT" {
			internal.PutWant("/topic"+path, jsonStr)
		} else {
			internal.DeleteWant("/topic"+path, jsonStr)
		}
	}
}

type topicGroupnameJSON struct {
	Topic     string `json:"topic"`
	Groupname string `json:"groupname"`
	Recursive bool   `json:"recursive"`
}

func topicAddDeleteGroups(method string, topic string, groups []string, path string) {
	for _, groupname := range groups {
		t := topicGroupnameJSON{topic, groupname, recursive}
		jsonStr, err := json.Marshal(t)
		internal.Check(err)
		if method == "PUT" {
			internal.PutWant("/topic"+path, jsonStr)
		} else {
			internal.DeleteWant("/topic"+path, jsonStr)
		}
	}
}
