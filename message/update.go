package message

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cmdMessageUpdate = &cobra.Command{
	Use:     "update",
	Aliases: []string{"up"},
	Short:   "Update a message (if it's enabled on topic): tatcli message update <topic> <idMessage> <my message...>",
	Long: `Update a message:
	tatcli message update <topic> <idMessage> <my message...>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 3 {
			topic := args[0]
			messageID := args[1]
			message := strings.Join(args[2:len(args)], " ")
			messageAction("update", topic, messageID, message, "")
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument to update a message: tatcli message update --help\n")
			cmd.Help()
		}
	},
}
