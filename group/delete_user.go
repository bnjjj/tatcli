package group

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmdGroupDeleteUser = &cobra.Command{
	Use:   "deleteUser",
	Short: "Delete Users from a group: tacli group deleteUser <groupname> <username1> [<username2> ... ]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			groupDeleteUsers(args[0], args[1:len(args)])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli group deleteUser --help\n")
		}
	},
}

func groupDeleteUsers(groupname string, users []string) {
	groupAddDeleteUsers("PUT", groupname, users, "/remove/user")
}
