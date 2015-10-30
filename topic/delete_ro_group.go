package topic

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicDeleteRoGroup.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Delete Rights RO recursively")
}

var cmdTopicDeleteRoGroup = &cobra.Command{
	Use:   "deleteRoGroup",
	Short: "Delete Read Only Groups from a topic: tatcli topic deleteRoGroup [--recursive] <topic> <groupname1> [<groupname2>]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicDeleteRoGroups(args[0], args[1:len(args)])
		} else {
			fmt.Println("Invalid argument: tatcli topic deleteRoGroup --help")
		}
	},
}

func topicDeleteRoGroups(topic string, groups []string) {
	topicAddDeleteGroups("DELETE", topic, groups, "/remove/rogroup")
}
