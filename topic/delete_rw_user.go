package topic

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicDeleteRwUser.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Delete Rights RW recursively")
}

var cmdTopicDeleteRwUser = &cobra.Command{
	Use:   "deleteRwUser",
	Short: "Delete Read Write Users from a topic: tatcli topic deleteRwUser [--recursive] <topic> <username1> [username2]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicDeleteRwUsers(args[0], args[1:len(args)])
		} else {
			fmt.Println("Invalid argument: tatcli topic deleteRwUser --help")
		}
	},
}

func topicDeleteRwUsers(topic string, users []string) {
	topicAddDeleteUsers("PUT", topic, users, "/remove/rwuser")
}
