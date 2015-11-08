package topic

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

func init() {
	cmdTopicDeleteParameter.Flags().BoolVarP(&recursive, "recursive", "r", false, "Remove Parameter recursively")
}

var cmdTopicDeleteParameter = &cobra.Command{
	Use:   "deleteParameter",
	Short: "Remove Parameter to a topic: tatcli topic deleteParameter [--recursive] <topic> <key> [<key2>]... ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicDeleteParameter(args[0], args[1:len(args)])
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument: tatcli topic deleteParameter --help\n")
		}
	},
}

func topicDeleteParameter(topic string, keys []string) {
	for _, key := range keys {
		t := topicParameterJSON{Topic: topic, Key: key, Recursive: recursive}
		jsonStr, err := json.Marshal(t)
		internal.Check(err)
		internal.PutWant("/topic/remove/parameter", jsonStr)
	}
}
