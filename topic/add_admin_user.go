package topic

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicAddAdminUser.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Rights Admin recursively")
}

var cmdTopicAddAdminUser = &cobra.Command{
	Use:   "addAdminUser",
	Short: "Add Admin Users to a topic: tatcli topic addAdminUser [--recursive] <topic> <username1> [username2]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicAddAdminUsers(args[0], args[1:len(args)])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli topic addAdminUser --help\n")
		}
	},
}

func topicAddAdminUsers(topic string, users []string) {
	topicAddDeleteUsers("PUT", topic, users, "/add/adminuser")
}
