package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdUserAddFavoriteTopic = &cobra.Command{
	Use:   "addFavoriteTopic",
	Short: "Add a favorite Topic: tatcli user addFavoriteTopic <topicName>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			userPOST("/me/topics" + args[0])
		} else {
			fmt.Println("Invalid argument: tatcli user addFavoriteTopic --help")
		}
	},
}
