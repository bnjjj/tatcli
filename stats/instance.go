package stats

import (
	"fmt"
	"os"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdStatsInstance = &cobra.Command{
	Use:   "instance",
	Short: "Info about current instance of engine",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli stats instance --help\n")
			cmd.Usage()
		} else {
			fmt.Print(internal.GetWantReturn("/stats/instance"))
		}
	},
}
