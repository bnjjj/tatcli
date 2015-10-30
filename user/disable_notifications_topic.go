package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdUserDisableNotificationsTopic = &cobra.Command{
	Use:   "disableNotificationsTopic",
	Short: "Disable notifications on a topic: tatcli user disableNotificationsTopic <topicName>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			userPOST("/me/disable/notifications/topics" + args[0])
		} else {
			fmt.Println("Invalid argument: tatcli user disableNotificationsTopic --help")
		}
	},
}
