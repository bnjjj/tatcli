package stats

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsDBReplSetGetStatus = &cobra.Command{
	Use:   "dbReplSetGetStatus",
	Short: "DB Stats: tatcli stats dbReplSetGetStatus",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Println("Invalid argument: tatcli stats dbReplSetGetStatus --help")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/db/replSetGetStatus"))
		}
	},
}
