package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdUserAddFavoriteTag = &cobra.Command{
	Use:   "addFavoriteTag",
	Short: "Add a favorite Tag: tatcli user addFavoriteTag <tag>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			userPOST("/me/tags/" + args[0])
		} else {
			fmt.Println("Invalid argument: tatcli user addFavoriteTag --help")
		}
	},
}
