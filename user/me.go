package user

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdUserMe = &cobra.Command{
	Use:   "me",
	Short: "Get Information about you: tatcli user me",
	Run: func(cmd *cobra.Command, args []string) {
		user("/me")
	},
}

func user(path string) {
	fmt.Print(internal.GetWantReturn("/user" + path))
}
