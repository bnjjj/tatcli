package message

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdMessageDelete = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"unbookmark", "rm"},
	Short:   "Remove a message (or bookmark) from Private Topic: tatcli message delete <idMessage>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			messageAction("remove", "", args[0], "", "")
		} else {
			fmt.Println("Invalid argument to delete message: tatcli message delete --help")
		}
	},
}
