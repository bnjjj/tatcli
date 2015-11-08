package topic

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicAddAdminGroup.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Rights Admin recursively")
}

var cmdTopicAddAdminGroup = &cobra.Command{
	Use:   "addAdminGroup",
	Short: "Add Admin Groups to a topic: tatcli topic addAdminGroup [--recursive] <topic> <groupname1> [groupname2]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicAddAdminGroups(args[0], args[1:len(args)])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli topic addAdminGroup --help\n")
		}
	},
}

func topicAddAdminGroups(topic string, groups []string) {
	topicAddDeleteGroups("PUT", topic, groups, "/add/admingroup")
}
