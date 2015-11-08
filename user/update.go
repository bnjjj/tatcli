package user

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdUserUpdate = &cobra.Command{
	Use:   "update",
	Short: "Update Fullname and Email of a user (admin only): tatcli user update <username> <newEmail> <newFullname>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 3 {
			username := args[0]
			email := args[1]
			fullname := strings.Join(args[2:len(args)], " ")
			updateUserAction(username, email, fullname)
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli user update --help\n")
		}
	},
}

type updateUserJSON struct {
	Username    string `json:"username"`
	NewEmail    string `json:"newEmail"`
	NewFullname string `json:"newFullname"`
}

func updateUserAction(username, newEmail, newFullname string) {
	m := updateUserJSON{username, newEmail, newFullname}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	internal.PutWant("/user/update", jsonStr)
}
