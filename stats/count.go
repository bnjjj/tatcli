package stats

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsCount = &cobra.Command{
	Use:   "count",
	Short: "Count all messages, groups, presences, users, groups, topics: tatcli stats count",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Println("Invalid argument: tatcli stats count --help")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/count"))
		}
	},
}
