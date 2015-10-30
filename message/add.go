package message

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

func init() {
	cmdMessageAdd.Flags().IntVarP(&dateCreation, "dateCreation", "", -1, "Force date creation, only for system user")
}

var cmdMessageAdd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "tatcli message add [--dateCreation=timestamp] <topic> <my message>",
	Long: `Add a message to a Topic:
		tatcli message add /Private/firstname.lastname my new messsage
		`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topic := args[0]
			message := strings.Join(args[1:len(args)], " ")
			messageCreate(topic, message)
		} else {
			fmt.Println("Invalid argument to add a message: tatcli msg add --help")
		}
	},
}

type messageJSON struct {
	Text         string `json:"text"`
	DateCreation int    `json:"dateCreation,omitempty"`
}

func messageCreate(topic, message string) {
	m := messageJSON{Text: message}
	if dateCreation > 0 {
		m.DateCreation = dateCreation
	}
	jsonStr, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Print(internal.PostWant("/message"+topic, jsonStr))
}
