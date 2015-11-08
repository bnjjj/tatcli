package stats

import (
	"fmt"
	"os"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsCount = &cobra.Command{
	Use:   "count",
	Short: "Count all messages, groups, presences, users, groups, topics: tatcli stats count",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli stats count --help\n")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/count"))
		}
	},
}
