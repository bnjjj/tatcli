package stats

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsDBServerStatus = &cobra.Command{
	Use:   "dbServerStatus",
	Short: "DB Stats: tatcli stats dbServerStatus",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Println("Invalid argument: tatcli stats dbServerStatus --help")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/db/serverStatus"))
		}
	},
}
