package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdUserResetSystem = &cobra.Command{
	Use:   "resetSystemUser",
	Short: "Reset password for a system user (admin only): tatcli user resetSystemUser <username>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			resetSystemUserAction(args[0])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli user resetSystemUser --help\n")
		}
	},
}

type resetSystemUserJSON struct {
	Username string `json:"username"`
}

func resetSystemUserAction(username string) {
	m := resetSystemUserJSON{username}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	fmt.Print(internal.ReqWant("PUT", http.StatusCreated, "/user/resetsystem", jsonStr))
}
