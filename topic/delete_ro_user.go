package topic

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicDeleteRoUser.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Delete Rights RO recursively")
}

var cmdTopicDeleteRoUser = &cobra.Command{
	Use:   "deleteRoUser",
	Short: "Delete Read Only Users from a topic: tatcli topic deleteRoUser [--recursive] <topic> <username1> [username2]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicDeleteRoUsers(args[0], args[1:len(args)])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli topic deleteRoUser --help\n")
		}
	},
}

func topicDeleteRoUsers(topic string, users []string) {
	topicAddDeleteUsers("PUT", topic, users, "/remove/rouser")
}
