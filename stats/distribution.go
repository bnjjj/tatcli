package stats

import (
	"fmt"
	"os"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsDistribution = &cobra.Command{
	Use:   "distribution",
	Short: "Distribution of messages per topics: tatcli stats distribution",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli stats distribution --help\n")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/distribution"))
		}
	},
}
