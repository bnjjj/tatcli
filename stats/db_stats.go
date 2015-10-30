package stats

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsDBStats = &cobra.Command{
	Use:   "dbstats",
	Short: "DB Stats: tatcli stats dbstats",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Println("Invalid argument: tatcli stats db --help")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/db/stats"))
		}
	},
}
