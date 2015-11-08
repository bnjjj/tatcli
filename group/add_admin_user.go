package group

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdGroupAddAdminUser = &cobra.Command{
	Use:   "addAdminUser",
	Short: "Add Admin Users to a group: tacli group addAdminUser <groupname> <username1> [<username2> ... ]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			groupAddAdminUsers(args[0], args[1:len(args)])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli group addAdminUser --help\n")
		}
	},
}

func groupAddAdminUsers(groupname string, users []string) {
	groupAddDeleteUsers("PUT", groupname, users, "/add/adminuser")
}

type groupUsernameJSON struct {
	Groupname string `json:"groupname"`
	Username  string `json:"username"`
}

func groupAddDeleteUsers(method string, groupname string, users []string, path string) {
	for _, username := range users {
		t := groupUsernameJSON{groupname, username}
		jsonStr, err := json.Marshal(t)
		internal.Check(err)
		if method == "PUT" {
			internal.PutWant("/group"+path, jsonStr)
		} else {
			internal.DeleteWant("/group"+path, jsonStr)
		}
	}
}
