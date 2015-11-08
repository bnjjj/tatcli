package stats

import (
	"fmt"
	"os"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsDBReplSetGetConfig = &cobra.Command{
	Use:   "dbReplSetGetConfig",
	Short: "DB Stats: tatcli stats dbReplSetGetConfig",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli stats dbReplSetGetConfig --help\n")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/db/replSetGetConfig"))
		}
	},
}
