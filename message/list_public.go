package message

import (
	"fmt"
	"os"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdMessageListPublic = &cobra.Command{
	Use:     "read",
	Short:   "List all messages on one public topic (read only): tatcli msg read <Topic> <skip> <limit>",
	Aliases: []string{"lp"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			skip, limit := internal.GetSkipLimit(args)
			messagesList(args[0], skip, limit)
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument to list message: See tatcli msg list --help\n")
		}
	},
}
