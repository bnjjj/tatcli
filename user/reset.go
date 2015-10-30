package user

import (
	"encoding/json"
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdUserReset = &cobra.Command{
	Use:   "reset",
	Short: "Ask for Reset a password: tatcli user reset <username> <email>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			userReset(args[0], args[1])
		} else {
			fmt.Println("Invalid argument to reset password: tatcli user reset --help")
		}
	},
}

func userReset(username, email string) {
	ssl := ""
	if internal.SSLInsecureSkipVerify {
		ssl = "--sslInsecureSkipVerify=true"
	}

	m := userJSON{
		Username: username,
		Email:    email,
		Callback: fmt.Sprintf("tatcli %s --url=:scheme://:host::port:path user verify --save :username :token", ssl),
	}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	fmt.Printf(internal.PostWant("/user/reset", jsonStr))
}
