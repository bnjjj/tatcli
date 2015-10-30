package user

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var withGroups bool

func init() {
	cmdUserList.Flags().BoolVarP(&withGroups, "with-groups", "g", false, "List Users with groups, admin only")
}

var cmdUserList = &cobra.Command{
	Use:   "list",
	Short: "List all users: tatcli user list [<skip>] [<limit>]",
	Run: func(cmd *cobra.Command, args []string) {
		skip, limit := internal.GetSkipLimit(args)
		usersList(skip, limit)
	},
}

func usersList(skip, limit string) {
	if withGroups {
		fmt.Print(internal.GetWantReturn(fmt.Sprintf("/users?skip=%s&limit=%s&withGroups=true", skip, limit)))
	} else {
		fmt.Print(internal.GetWantReturn(fmt.Sprintf("/users?skip=%s&limit=%s", skip, limit)))
	}
}
