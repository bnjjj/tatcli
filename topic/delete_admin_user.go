package topic

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicDeleteAdminUser.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Delete Rights Admin recursively")
}

var cmdTopicDeleteAdminUser = &cobra.Command{
	Use:   "deleteAdminUser",
	Short: "Delete Admin Users from a topic: tatcli topic deleteAdminUser [--recursive] <topic> <username1> [username2]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicDeleteAdminUsers(args[0], args[1:len(args)])
		} else {
			fmt.Println("Invalid argument: tatcli topic deleteAdminUser --help")
		}
	},
}

func topicDeleteAdminUsers(topic string, users []string) {
	topicAddDeleteUsers("PUT", topic, users, "/remove/adminuser")
}
