package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdUserRemoveFavoriteTag = &cobra.Command{
	Use:   "removeFavoriteTag",
	Short: "Remove a favorite Tag: tatcli user removeFavoriteTag <tag>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			userDelete("/me/tags/" + args[0])
		} else {
			fmt.Println("Invalid argument: tatcli user removeFavoriteTag --help")
		}
	},
}
