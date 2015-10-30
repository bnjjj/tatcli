package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdUserAddContact = &cobra.Command{
	Use:   "addContact",
	Short: "Add a contact: tatcli user addContact <contactUsername>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			userPOST("/me/contacts/" + args[0])
		} else {
			fmt.Println("Invalid argument to add contact: tatcli user addContact --help")
		}
	},
}
