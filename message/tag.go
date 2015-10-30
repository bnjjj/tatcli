package message

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdMessageTag = &cobra.Command{
	Use:   "tag",
	Short: "Add a tag to a message (user system with rights only): tatcli message tag <idMessage> <my Tag>",
	Long: `Add a tag to a message (user system with rights only):
	tatcli message tag <idMessage> <my tag>
	Example in bash:
	tatcli message tag idMessage myTag
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			tag := args[1]
			messageAction("tag", "/", args[0], tag, "")
		} else {
			fmt.Println("Invalid argument to add a tag: tatcli message tag --help")
		}
	},
}
