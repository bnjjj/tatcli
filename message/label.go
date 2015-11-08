package message

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cmdMessageLabel = &cobra.Command{
	Use:   "label",
	Short: "Add a label to a message: tatcli message label <idMessage> <colorInHexa> <my Label>",
	Long: `Add a label to a message:
	tatcli message label <Topic> <inReplyOfId> <my message...>
	Example in bash:
	tatcli message label /Internal/MyTopic aazzerrr \#EEEEEE my White Label
	or works too:
	tatcli message label /Internal/MyTopic aazzerrr EEEEEE my White Label
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 3 {
			text := strings.Join(args[2:len(args)], " ")
			color := args[1]
			if !strings.HasPrefix(color, "#") {
				color = "#" + color
			}
			messageAction("label", "/", args[0], text, color)
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument to add a label: tatcli message label --help\n")
		}
	},
}
