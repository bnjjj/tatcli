package topic

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

func init() {
	cmdTopicAddParameter.Flags().BoolVarP(&recursive, "recursive", "r", false, "Add Parameter recursively")
}

var cmdTopicAddParameter = &cobra.Command{
	Use:   "addParameter",
	Short: "Add Parameter to a topic: tatcli topic addParameter [--recursive] <topic> <key>:<value> [<key2>:<value2>]... ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			topicAddParameter(args[0], args[1:len(args)])
		} else {
			fmt.Println("Invalid argument: tatcli topic addParameter --help")
		}
	},
}

type topicParameterJSON struct {
	Topic     string `json:"topic"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Recursive bool   `json:"recursive"`
}

func topicAddParameter(topic string, parameters []string) {
	for _, param := range parameters {
		parameterSplitted := strings.Split(param, ":")
		if len(parameterSplitted) != 2 {
			continue
		}
		t := topicParameterJSON{Topic: topic, Key: parameterSplitted[0], Value: parameterSplitted[1], Recursive: recursive}
		jsonStr, err := json.Marshal(t)
		internal.Check(err)
		internal.PutWant("/topic/add/parameter", jsonStr)
	}
}
