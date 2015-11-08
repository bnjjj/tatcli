package user

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmdUserRemoveContact = &cobra.Command{
	Use:   "removeContact",
	Short: "Remove a contact: tatcli user removeContact <contactUsername>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			userDelete("/me/contacts/" + args[0])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli user removeContact --help\n")
		}
	},
}
