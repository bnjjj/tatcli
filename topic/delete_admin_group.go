package topic

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicDeleteAdminGroup.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Delete Rights Admin recursively")
}

var cmdTopicDeleteAdminGroup = &cobra.Command{
	Use:   "deleteAdminGroup",
	Short: "Delete Admin Groups from a topic: tatcli topic deleteAdminGroup [--recursive] <topic> <groupname1> [<groupname2>]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicDeleteAdminGroups(args[0], args[1:len(args)])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli topic deleteAdminGroup --help\n")
		}
	},
}

func topicDeleteAdminGroups(topic string, groups []string) {
	topicAddDeleteGroups("PUT", topic, groups, "/remove/admingroup")
}
