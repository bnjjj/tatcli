package message

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cmdMessageUntask = &cobra.Command{
	Use:   "untask",
	Short: "Remove a message from tasks: tatcli message untask /Private/username/tasks idMessage",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			idMessage := strings.Join(args[1:len(args)], " ")
			messageAction("untask", args[0], idMessage, "", "")
		} else {
			fmt.Println("Invalid argument to untask a message: tatcli message untask --help")
		}
	},
}
