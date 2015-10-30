package message

import (
	"encoding/json"
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cmdMessageAdd)
	Cmd.AddCommand(cmdMessageReply)
	Cmd.AddCommand(cmdMessageBookmark)
	Cmd.AddCommand(cmdMessageDelete)
	Cmd.AddCommand(cmdMessageUpdate)
	Cmd.AddCommand(cmdMessageTask)
	Cmd.AddCommand(cmdMessageUntask)
	Cmd.AddCommand(cmdMessageLike)
	Cmd.AddCommand(cmdMessageUnlike)
	Cmd.AddCommand(cmdMessageLabel)
	Cmd.AddCommand(cmdMessageUnlabel)
	Cmd.AddCommand(cmdMessageTag)
	Cmd.AddCommand(cmdMessageUntag)
	Cmd.AddCommand(cmdMessageList)
	Cmd.AddCommand(cmdMessageListPublic)
}

// Cmd message
var Cmd = &cobra.Command{
	Use:     "message",
	Short:   "Manipulate messages: tatcli message --help",
	Long:    `Manipulate messages: tatcli message <command>`,
	Aliases: []string{"m", "msg"},
}

type messageActionJSON struct {
	Text        string `json:"text"`
	IDReference string `json:"idReference"`
	Action      string `json:"action"`
	Option      string `json:"option"`
}

func messageAction(action, topic, idReference, message, option string) {
	m := messageActionJSON{message, idReference, action, option}
	jsonStr, err := json.Marshal(m)
	internal.Check(err)
	if action == "remove" {
		internal.DeleteWant("/message/"+idReference, nil)
	} else if action == "like" || action == "unlike" ||
		action == "label" || action == "unlabel" ||
		action == "task" || action == "untask" ||
		action == "tag" || action == "untag" ||
		action == "update" {
		internal.PutWant("/message"+topic, jsonStr)
	} else {
		fmt.Print(internal.PostWant(fmt.Sprintf("/message%s", topic), jsonStr))
	}
}
