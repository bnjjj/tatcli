package message

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmdMessageLike = &cobra.Command{
	Use:   "like",
	Short: "Like a message: tatcli message like <idMessage>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			messageAction("like", "/", args[0], "", "")
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument to like a message: tatcli message like --help\n")
		}
	},
}
