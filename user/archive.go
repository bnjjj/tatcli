package user

import (
	"encoding/json"
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdUserArchive = &cobra.Command{
	Use:   "archive",
	Short: "Archive a user (admin only): tatcli user archive <username>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			archiveUserAction(args[0])
		} else {
			fmt.Println("Invalid argument: tatcli user archive --help")
		}
	},
}

type usernameUserJSON struct {
	Username string `json:"username"`
}

func archiveUserAction(username string) {
	m := usernameUserJSON{username}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	internal.PutWant("/user/archive", jsonStr)
}
