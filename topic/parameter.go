package topic

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

func init() {
	cmdTopicParameter.Flags().BoolVarP(&recursive, "recursive", "r", false, "Update param topic recursively")
}

var cmdTopicParameter = &cobra.Command{
	Use:     "parameter",
	Short:   "Update param on one topic: tatcli topic param [--recursive] <topic> <maxLength> <canForceDate> <canUpdateMsg> <canDeleteMsg> <canUpdateAllMsg> <canDeleteAllMsg> <isROPublic>",
	Aliases: []string{"param"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 8 {
			fmt.Fprintf(os.Stderr, "Invalid parameter to tatcli topic param. See tatcli topic param --help\n")
			cmd.Help()
			os.Exit(1)
		}
		maxLength, err := strconv.Atoi(args[1])
		internal.Check(err)
		canForceDate, err := strconv.ParseBool(args[2])
		internal.Check(err)
		canUpdateMsg, err := strconv.ParseBool(args[3])
		internal.Check(err)
		canDeleteMsg, err := strconv.ParseBool(args[4])
		internal.Check(err)
		canUpdateAllMsg, err := strconv.ParseBool(args[5])
		internal.Check(err)
		canDeleteAllMsg, err := strconv.ParseBool(args[6])
		internal.Check(err)
		isROPublic, err := strconv.ParseBool(args[7])
		internal.Check(err)
		topicParam(args[0], maxLength, canForceDate, canUpdateMsg, canDeleteMsg, canUpdateAllMsg, canDeleteAllMsg, isROPublic)
	},
}

type paramJSON struct {
	Topic           string `json:"topic"`
	MaxLength       int    `json:"maxlength"`
	CanForceDate    bool   `json:"canForceDate"`
	CanUpdateMsg    bool   `json:"canUpdateMsg"`
	CanDeleteMsg    bool   `json:"canDeleteMsg"`
	CanUpdateAllMsg bool   `json:"canUpdateAllMsg"`
	CanDeleteAllMsg bool   `json:"canDeleteAllMsg"`
	IsROPublic      bool   `json:"isROPublic"`
	Recursive       bool   `json:"recursive"`
}

func topicParam(topic string, maxLength int, canForceDate, canUpdateMsg, canDeleteMsg, canUpdateAllMsg, canDeleteAllMsg, isROPublic bool) {
	t := paramJSON{
		Topic:           topic,
		MaxLength:       maxLength,
		CanForceDate:    canForceDate,
		CanUpdateMsg:    canUpdateMsg,
		CanDeleteMsg:    canDeleteMsg,
		CanUpdateAllMsg: canUpdateAllMsg,
		CanDeleteAllMsg: canDeleteAllMsg,
		IsROPublic:      isROPublic,
		Recursive:       recursive,
	}
	jsonStr, err := json.Marshal(t)
	internal.Check(err)
	internal.PutWant("/topic/param", jsonStr)
}
