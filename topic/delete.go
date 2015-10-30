package topic

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdTopicDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a new topic: tatcli delete <topic>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			topicDelete(args[0])
		} else {
			fmt.Println("Invalid argument: tatcli topic delete --help")
		}
	},
}

func topicDelete(topic string) {
	internal.DeleteWant("/topic"+topic, nil)
}
