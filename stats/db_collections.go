package stats

import (
	"fmt"
	"os"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsDBCollections = &cobra.Command{
	Use:   "dbCollections",
	Short: "DB Stats on each collection: tatcli stats dbCollections",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli stats dbCollections --help\n")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/db/collections"))
		}
	},
}
