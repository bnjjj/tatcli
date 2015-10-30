package presence

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdPresenceAdd = &cobra.Command{
	Use:   "add",
	Short: "Add a new presence on one topic with status (online, offline, busy): tatcli presence add <topic> <status>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			status := strings.Join(args[1:len(args)], " ")
			presenceCreate(args[0], status)
		} else {
			fmt.Println("Invalid argument: tatcli presence add --help")
		}
	},
}

type presenceJSON struct {
	Status string `json:"status"`
}

func presenceCreate(topic, status string) {
	p := presenceJSON{status}
	jsonStr, err := json.Marshal(p)
	internal.Check(err)
	internal.PostWant("/presenceget"+topic, jsonStr)
}
