package group

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdGroupAddUser = &cobra.Command{
	Use:   "addUser",
	Short: "Add Users to a group: tacli group addUser <groupname> <username1> [<username2> ... ]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			groupAddUsers(args[0], args[1:len(args)])
		} else {
			fmt.Println("Invalid argument: tatcli group addUser --help")
		}
	},
}

func groupAddUsers(groupname string, users []string) {
	groupAddDeleteUsers("PUT", groupname, users, "/add/user")
}
