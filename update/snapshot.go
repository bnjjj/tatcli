package update

import "github.com/spf13/cobra"

// used by CI to inject url for downloading with tatcli update.
// value of urlUpdate injected at build time
// full URL update is constructed with architecture var :
// urlUpdate + "tatcli-" + architecture, tatcli is the binary
var urlUpdateSnapshot string

var cmdUpdateSnapshot = &cobra.Command{
	Use:     "snapshot",
	Short:   "Update tatcli to latest snapshot version: tatcli update snapshot",
	Long:    `tatcli update snapshot`,
	Aliases: []string{"snap"},
	Run: func(cmd *cobra.Command, args []string) {
		doUpdate(urlUpdateSnapshot, architecture)
	},
}
