package topic

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdTopicList = &cobra.Command{
	Use:     "list",
	Short:   "List all topics: tatcli topic list [<skip>] [<limit>] [<true>] if true, return unread cound msg",
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		skip := "0"
		limit := "10"
		if len(args) >= 2 {
			skip = args[0]
			limit = args[1]
		}

		getNbMsgUnread := false
		if len(args) == 3 && args[2] == "true" {
			getNbMsgUnread = true
		}
		topicsList(skip, limit, getNbMsgUnread)
	},
}

func topicsList(skip string, limit string, getNbMsgUnread bool) {
	n := ""
	if getNbMsgUnread {
		n = "&getNbMsgUnread=true"
	}
	fmt.Print(internal.GetWantReturn(fmt.Sprintf("/topics?skip=%s&limit=%s%s", skip, limit, n)))
}
