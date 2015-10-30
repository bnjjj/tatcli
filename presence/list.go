package presence

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdPresenceList = &cobra.Command{
	Use:   "list",
	Short: "List all presences on one topic: tatcli presence list <topic> [<skip>] [<limit>]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			skip, limit := internal.GetSkipLimit(args)
			presencesList(args[0], skip, limit)
		} else {
			fmt.Println("Invalid argument: tatcli presence list --help")
		}
	},
}

func presencesList(topic string, skip, limit string) {
	fmt.Print(internal.GetWantReturn(fmt.Sprintf("/presences%s?skip=%s&limit=%s", topic, skip, limit)))
}
