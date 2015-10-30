package message

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cmdMessageUnlabel = &cobra.Command{
	Use:   "unlabel",
	Short: "Remove a label from a message: tatcli message unlabel <idMessage> <my Label>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			label := strings.Join(args[1:len(args)], " ")
			messageAction("unlabel", "/", args[0], label, "")
		} else {
			fmt.Println("Invalid argument to unlabel a message: tatcli message unlabel --help")
		}
	},
}
