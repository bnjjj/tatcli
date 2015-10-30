package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdUserEnableNotificationsTopic = &cobra.Command{
	Use:   "enableNotificationsTopic",
	Short: "Enable notifications on a topic: tatcli user enableNotificationsTopic <topicName>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			userPOST("/me/enable/notifications/topics" + args[0])
		} else {
			fmt.Println("Invalid argument: tatcli user enableNotificationsTopic --help")
		}
	},
}
