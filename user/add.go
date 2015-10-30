package user

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdUserAdd = &cobra.Command{
	Use:   "add",
	Short: "Add a user: tatcli user add <username> <email> <fullname>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 3 {
			username := args[0]
			email := args[1]
			fullname := strings.Join(args[2:len(args)], " ")
			userAdd(username, email, fullname)
		} else {
			fmt.Println("Invalid argument to add user: tatcli user add --help")
		}
	},
}

type userJSON struct {
	Username string `json:"username"`
	Fullname string `json:"fullname,omitempty"`
	Email    string `json:"email"`
	// Callback contains command to execute to verify account
	// this command is displayed in ask for confirmation mail
	Callback string `json:"callback"`
}

func userAdd(username, email, fullname string) {
	m := userJSON{username, fullname, email, "tatcli --url=:scheme://:host::port:path user verify --save :username :token"}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	fmt.Printf(internal.PostWant("/user", jsonStr))
}
