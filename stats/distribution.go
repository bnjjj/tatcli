package stats

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsDistribution = &cobra.Command{
	Use:   "distribution",
	Short: "Distribution of messages per topics: tatcli stats distribution",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Println("Invalid argument: tatcli stats distribution --help")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/distribution"))
		}
	},
}
