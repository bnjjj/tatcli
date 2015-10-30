package group

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdGroupList = &cobra.Command{
	Use:   "list",
	Short: "List all groups: tatcli group list <skip> <limit>",
	Run: func(cmd *cobra.Command, args []string) {
		skip, limit := internal.GetSkipLimit(args)
		groupsList(skip, limit)
	},
}

func groupsList(skip, limit string) {
	fmt.Print(internal.GetWantReturn(fmt.Sprintf("/groups?skip=%s&limit=%s", skip, limit)))
}
