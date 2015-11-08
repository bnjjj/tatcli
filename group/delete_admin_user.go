package group

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmdGroupDeleteAdminUser = &cobra.Command{
	Use:   "deleteAdminUser",
	Short: "Delete Admin Users from a group: tacli group deleteAdminUser <groupname> <username1> [<username2> ... ]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			groupDeleteAdminUsers(args[0], args[1:len(args)])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli group deleteAdminUser --help\n")
		}
	},
}

func groupDeleteAdminUsers(groupname string, users []string) {
	groupAddDeleteUsers("PUT", groupname, users, "/remove/adminuser")
}
