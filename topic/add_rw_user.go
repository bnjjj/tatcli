package topic

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicAddRwUser.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Rights RW recursively")
}

var cmdTopicAddRwUser = &cobra.Command{
	Use:   "addRwUser",
	Short: "Add Read Write Users to a topic: tatcli topic addRwUser [--recursive] <topic> <username1> [username2]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicAddRwUsers(args[0], args[1:len(args)])
		} else {
			fmt.Println("Invalid argument: tatcli topic addRwUser --help")
		}
	},
}

func topicAddRwUsers(topic string, users []string) {
	topicAddDeleteUsers("PUT", topic, users, "/add/rwuser")
}
