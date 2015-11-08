package message

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmdMessageUntag = &cobra.Command{
	Use:   "untag",
	Short: "Remove a tag from a message (user system with rights only): tatcli message untag <idMessage> <myTag>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			tag := args[1]
			messageAction("untag", "/", args[0], tag, "")
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument to untag a message: tatcli message untag --help\n")
		}
	},
}
