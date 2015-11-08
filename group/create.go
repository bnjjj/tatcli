package group

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdGroupCreate = &cobra.Command{
	Use:   "create",
	Short: "create a new group: tatlic group create <groupname> <description>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			description := strings.Join(args[1:len(args)], " ")
			groupCreate(args[0], description)
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli group create --help\n")
		}
	},
}

type groupJSON struct {
	Groupname   string `json:"name"`
	Description string `json:"description"`
}

func groupCreate(group, description string) {
	m := groupJSON{group, description}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	internal.PostWant("/group", jsonStr)
}
