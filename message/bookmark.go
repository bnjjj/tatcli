package message

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cmdMessageBookmark = &cobra.Command{
	Use:   "bookmark",
	Short: "Bookmark a message to a topic: tatcli message bookmark /Private/username/bookmarks/sub-topic idMessage",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			topic := args[0]
			idMessage := strings.Join(args[1:len(args)], " ")
			messageAction("bookmark", topic, idMessage, "", "")
		} else {
			fmt.Println("Invalid argument to bookmark message: tatcli message bookmark --help")
		}
	},
}
