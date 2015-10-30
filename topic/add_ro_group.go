package topic

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicAddRoGroup.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Rights RO recursively")
}

var cmdTopicAddRoGroup = &cobra.Command{
	Use:   "addRoGroup",
	Short: "Add Read Only Groups to a topic: tatcli topic addRoGroup [--recursive] <topic> <groupname1> [<groupname2>]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicAddRoGroups(args[0], args[1:len(args)])
		} else {
			fmt.Println("Invalid argument: tatcli topic addRoGroup --help")
		}
	},
}

func topicAddRoGroups(topic string, groups []string) {
	topicAddDeleteGroups("PUT", topic, groups, "/add/rogroup")
}
