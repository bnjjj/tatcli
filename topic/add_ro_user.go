package topic

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicAddRoUser.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Rights RO recursively")
}

var cmdTopicAddRoUser = &cobra.Command{
	Use:   "addRoUser",
	Short: "Add Read Only Users to a topic: tatcli topic addRoUser [--recursive] <topic> <username1> [username2]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicAddRoUsers(args[0], args[1:len(args)])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli topic addRoUser --help\n")
		}
	},
}

func topicAddRoUsers(topic string, users []string) {
	topicAddDeleteUsers("PUT", topic, users, "/add/rouser")
}
