package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdUserConvertToSystem = &cobra.Command{
	Use:   "convert",
	Short: "Convert a user to a system user (admin only): tatcli user convert <username> <canWriteNotifications>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			convertUserAction(args[0], args[1])
		} else {
			fmt.Println("Invalid argument : tatcli user convert --help")
		}
	},
}

type convertUserJSON struct {
	Username              string `json:"username"`
	CanWriteNotifications string `json:"canWriteNotifications"`
}

func convertUserAction(username, canWriteNotifications string) {
	m := convertUserJSON{username, canWriteNotifications}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	fmt.Print(internal.ReqWant("PUT", http.StatusCreated, "/user/convert", jsonStr))
}
