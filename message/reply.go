package message

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cmdMessageReply = &cobra.Command{
	Use:     "reply",
	Aliases: []string{"r"},
	Short:   "Reply to a message: tatcli message reply <topic> <inReplyOfId> <my message...>",
	Long: `Reply to a message:
	tatcli message reply <topic> <inReplyOfId> <my message...>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 3 {
			topic := args[0]
			inReplyOfID := args[1]
			message := strings.Join(args[2:len(args)], " ")
			messageAction("reply", topic, inReplyOfID, message, "")
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument to reply to a message: tatcli message reply --help\n")
		}
	},
}
