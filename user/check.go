package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdUserCheck = &cobra.Command{
	Use:   "check",
	Short: "Check Private Topics and Default Group on one user (admin only): tatcli user check <username> <fixPrivateTopics> <fixDefaultGroup>",
	Long: `Check Private Topics and Default Group on one user:
tatcli user check <username> <fixPrivateTopics> <fixDefaultGroup>

Example :

tatcli check username true true
		`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 3 {
			fixPrivateTopics, e1 := strconv.ParseBool(args[1])
			fixDefaultGroup, e2 := strconv.ParseBool(args[2])
			if e1 != nil || e2 != nil {
				fmt.Println("Invalid argument: tatcli user check --help")
			}
			checkUserAction(args[0], fixPrivateTopics, fixDefaultGroup)
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli user check --help\n")
		}
	},
}

type checkUserJSON struct {
	Username         string `json:"username"`
	FixPrivateTopics bool   `json:"fixPrivateTopics"`
	FixDefaultGroup  bool   `json:"fixDefaultGroup"`
}

func checkUserAction(username string, fixPrivateTopics, fixDefaultGroup bool) {
	m := checkUserJSON{username, fixPrivateTopics, fixDefaultGroup}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	fmt.Println(internal.ReqWant("PUT", http.StatusCreated, "/user/check", jsonStr))
}
