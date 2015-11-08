package stats

import (
	"fmt"
	"os"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsDBServerStatus = &cobra.Command{
	Use:   "dbServerStatus",
	Short: "DB Stats: tatcli stats dbServerStatus",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli stats dbServerStatus --help\n")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/db/serverStatus"))
		}
	},
}
