package socket

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cmdSocketMessagesNew = &cobra.Command{
	Use:   "messagesNew",
	Short: "Open websocket and get events on new messages on one or many topics: tatcli socket messagesNew <<topic>|all>> [topic]...",
	Long: `Example:
	tatcli socket messagesNew all
	tatcli socket messagesNew /Topic/SubTopicA /Topic/SubTopicB
	`,
	Aliases: []string{"n", "new", "newMsg"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Invalid args. See tatcli socket messagesNew --help")
		} else {
			socketMessagesNew(strings.Join(args, " "))
		}
	},
}

func socketMessagesNew(args string) {
	c := newClient()
	wsActionSubscribeMessagesNew(c, args)
	done := make(chan bool)
	go socketRead(c)
	<-done
}
