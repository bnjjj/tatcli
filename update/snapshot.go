package update

import "github.com/spf13/cobra"

var cmdUpdateSnapshot = &cobra.Command{
	Use:     "snapshot",
	Short:   "Update tatcli to latest snapshot version: tatcli update snapshot",
	Long:    `tatcli update snapshot`,
	Aliases: []string{"snap"},
	Run: func(cmd *cobra.Command, args []string) {
		doUpdate(urlUpdateSnapshot, architecture)
	},
}
