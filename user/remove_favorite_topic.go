package user

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmdUserRemoveFavoriteTopic = &cobra.Command{
	Use:   "removeFavoriteTopic",
	Short: "Remove a favorite Topic: tatcli user removeFavoriteTopic <topicName>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			userDelete("/me/topics" + args[0])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli user removeFavoriteTopic --help\n")
		}
	},
}
