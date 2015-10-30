package user

import (
	"encoding/json"
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdUserSetAdmin = &cobra.Command{
	Use:   "setAdmin",
	Short: "Grant user to Tat admin (admin only): tatcli user setAdmin <username>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			setAdminUserAction(args[0])
		} else {
			fmt.Println("Invalid argument: tatcli user setAdmin --help")
		}
	},
}

func setAdminUserAction(username string) {
	m := usernameUserJSON{username}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	internal.PutWant("/user/setadmin", jsonStr)
}
