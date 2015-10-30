package stats

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsDBSlowestQueries = &cobra.Command{
	Use:   "dbSlowestQueries",
	Short: "DB Stats slowest Queries: tatcli stats dbSlowestQueries",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Println("Invalid argument: tatcli stats dbSlowestQueriess --help")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/db/slowestQueries"))
		}
	},
}
