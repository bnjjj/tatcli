package stats

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsDBCollections = &cobra.Command{
	Use:   "dbCollections",
	Short: "DB Stats on each collection: tatcli stats dbCollections",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Println("Invalid argument: tatcli stats dbCollections --help")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/db/collections"))
		}
	},
}
