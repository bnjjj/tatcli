package user

import (
	"encoding/json"
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdUserRename = &cobra.Command{
	Use:   "rename",
	Short: "Rename username of a user (admin only): tatcli user rename <oldUsername> <newUsername>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			renameUserAction(args[0], args[1])
		} else {
			fmt.Println("Invalid argument: tatcli user rename --help")
		}
	},
}

type renameUserJSON struct {
	Username    string `json:"username"`
	NewUsername string `json:"newUsername"`
}

func renameUserAction(username, newUsername string) {
	m := renameUserJSON{username, newUsername}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	internal.PutWant("/user/rename", jsonStr)
}
