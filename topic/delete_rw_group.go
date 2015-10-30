package topic

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicDeleteRwGroup.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Delete Rights RW recursively")
}

var cmdTopicDeleteRwGroup = &cobra.Command{
	Use:   "deleteRwGroup",
	Short: "Delete Read Write Groups from a topic: tatcli topic deleteRwGroup [--recursive] <topic> <groupname1> [<groupname2>]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicDeleteRwGroups(args[0], args[1:len(args)])
		} else {
			fmt.Println("Invalid argument: tatcli topic deleteRwGroup --help")
		}
	},
}

func topicDeleteRwGroups(topic string, groups []string) {
	topicAddDeleteGroups("DELETE", topic, groups, "/remove/rwgroup")
}
