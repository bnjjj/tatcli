package message

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmdMessageUnlike = &cobra.Command{
	Use:   "unlike",
	Short: "Unlike a message: tatcli message unlike <idMessage>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			messageAction("unlike", "/", args[0], "", "")
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument to unlike a message: tatcli message unlike --help\n")
		}
	},
}
