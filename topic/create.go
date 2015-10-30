package topic

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdTopicCreate = &cobra.Command{
	Use:   "create",
	Short: "Create a new topic: tatcli create <topic> <description of topic>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			description := strings.Join(args[1:len(args)], " ")
			topicCreate(args[0], description)
		} else {
			fmt.Println("Invalid argument: tatcli topic create --help")
		}
	},
}

type topicJSON struct {
	Topic       string `json:"topic"`
	Description string `json:"description"`
}

func topicCreate(topic, description string) {
	m := topicJSON{topic, description}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	internal.PostWant("/topic", jsonStr)
}
