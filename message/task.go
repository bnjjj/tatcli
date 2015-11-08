package message

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cmdMessageTask = &cobra.Command{
	Use:   "task",
	Short: "Create a task from one message to a topic: tatcli message task /Private/username/tasks/sub-topic idMessage",
	Long: `Create a task from one message to a topic:
	tatcli message task /Private/username/tasks/sub-topic idMessage`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			idMessage := strings.Join(args[1:len(args)], " ")
			messageAction("task", args[0], idMessage, "", "")
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument to task a message: tatcli message task --help\n")
		}
	},
}
