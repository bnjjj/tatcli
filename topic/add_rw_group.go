package topic

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cmdTopicAddRwGroup.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Rights RW recursively")
}

var cmdTopicAddRwGroup = &cobra.Command{
	Use:   "addRwGroup",
	Short: "Add Read Write Groups to a topic: tatcli topic addRwGroup [--recursive] <topic> <groupname1> [<groupname2>]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicAddRwGroups(args[0], args[1:len(args)])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli topic addRwGroup --help\n")
		}
	},
}

func topicAddRwGroups(topic string, groups []string) {
	topicAddDeleteGroups("PUT", topic, groups, "/add/rwgroup")
}
