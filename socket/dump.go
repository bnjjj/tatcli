package socket

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdSocketDump = &cobra.Command{
	Use:     "dump",
	Short:   "Dump websocket admin variables: tatcli socket dump",
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(internal.GetWantReturn("/sockets/dump"))
	},
}
