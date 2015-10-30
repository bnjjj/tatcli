package socket

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cmdSocketMessages = &cobra.Command{
	Use:   "messages",
	Short: "Open websocket and get events messages on one or many topics: tatcli socket messages <<topic>|all|onetree|fulltree>> [topic]...",
	Long: `Example:
	tatcli socket messages all
	tatcli socket messages /Topic/SubTopicA /Topic/SubTopicB
	tatcli socket messages fulltree all
	tatcli socket messages onetree all
	tatcli socket messages fulltree /Topic/SubTopicA /Topic/SubTopicB
	`,
	Aliases: []string{"m", "msg", "message"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Invalid args. See tatcli socket messages --help")
		} else {
			socketMessages(strings.Join(args, " "))
		}
	},
}

func socketMessages(args string) {
	c := newClient()
	wsActionSubscribeMessages(c, args)
	done := make(chan bool)
	go socketRead(c)
	<-done
}
