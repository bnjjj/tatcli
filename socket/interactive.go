package socket

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/ovh/tat/models"
	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdSocketInteractive = &cobra.Command{
	Use:   "interactive",
	Short: "Interactive mode Websocket: tatcli socket interactive (or tatcli socket i)",
	Long: `tatcli socket interactive

  You enter into an interactive mode. You can:
  - subscribe to events Messages on one topic, enter: subscribeMessages /myTopic
	- subscribe to events Messages on one topic, with one tree view, enter: subscribeMessages onetree /myTopic
	- subscribe to events Messages on one topic, with full tree view, enter: subscribeMessages fulltree /myTopic
  - unsubscribe to events Messages on one topic, enter: unsubscribeMessages /myTopic
  - subscribe to events Presences on one topic, enter: subscribePresences /myTopic
  - unsubscribe to events Presences on one topic, enter: unsubscribePresences /myTopic

  - subscribe to events Messages on all topics, enter: subscribeMessages all
	- subscribe to events Messages on all topics, with one tree view, enter: subscribeMessages onetree all
	- subscribe to events Messages on all topics, with full tree view, enter: subscribeMessages fulltree all
  - unsubscribe to events Messages on all topics, enter: unsubscribeMessages all
  - subscribe to events Presences on all topics, enter : subscribePresences all
  - unsubscribe to events Presences on all topics, enter: unsubscribePresences all

  - send a presence notification on one topic, enter: writePresence <status> <topic>
  Example: writePresence online /myTopic


	`,
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		socketInteractive()
	},
}

func socketRead(c *websocket.Conn) {
	for {
		_, r, err := c.ReadMessage()
		internal.Check(err)
		fmt.Print(color(string(r)))
	}
}
func socketInteractive() {
	c := newClient()

	done := make(chan bool)
	go socketRead(c)

	r := bufio.NewReader(os.Stdin)
	go func() {
		for {
			line, err := r.ReadString('\n')
			if err != nil && err.Error() != "unexpected newline" {
				internal.Check(err)
			}
			line = analyzeLine(c, line)
			if line != "" {
				if err = c.WriteMessage(1, []byte(line)); err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err.Error())
				}
			}
		}
	}()
	<-done
}

func color(msg string) string {
	return ("\033[36m" + msg + "\033[0m")
}

func analyzeLine(c *websocket.Conn, in string) string {
	line := strings.TrimSpace(in)

	s := strings.Split(line, " ")
	switch s[0] {
	case "subscribeMessages":
		wsActionSubscribeMessages(c, line[len("subscribeMessages"):])
	case "unsubscribeMessages":
		wsActionUnsubscribeMessages(c, line[len("unsubscribeMessages"):])
	case "subscribeMessagesNew":
		wsActionSubscribeMessagesNew(c, line[len("subscribeMessagesNew"):])
	case "unsubscribeMessagesNew":
		wsActionUnsubscribeMessagesNew(c, line[len("unsubscribeMessagesNew"):])
	case "subscribePresences":
		wsActionSubscribePresences(c, line[len("subscribePresences"):])
	case "unsubscribePresences":
		wsActionUnsubscribePresences(c, line[len("unsubscribePresences"):])
	case "subscribeUsers":
		wsActionSubscribeUsers(c)
	case "unsubscribeUsers":
		wsActionSubscribeUsers(c)
	case "writePresence":
		wsActionWritePresence(c, line[len("writePresence"):])
	default:
		return line
	}
	return ""
}

func getWSJSONTopic(action, args string) models.WSJSON {

	r := models.WSJSON{
		Action: "subscribeMessages",
	}

	for _, v := range strings.Split(strings.Trim(args, " "), " ") {
		if strings.HasPrefix(v, "/") || v == "all" {
			r.Topics = append(r.Topics, v)
		} else if v == "online" || v == "offline" || v == "busy" {
			r.Status = v
		} else if v == "onetree" || v == "fulltree" {
			r.TreeView = v
		}
	}
	return r
}

func getStrWSJSONTopic(w models.WSJSON) []byte {
	jsonStr, err := json.Marshal(w)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while converting to json: %s\n", err.Error())
	}
	return jsonStr
}

func wsActionSubscribeMessages(c *websocket.Conn, args string) {
	wsWrite(c, getStrWSJSONTopic(getWSJSONTopic("subscribeMessages", args)))
}

func wsActionUnsubscribeMessages(c *websocket.Conn, args string) {
	wsWrite(c, getStrWSJSONTopic(getWSJSONTopic("unsubscribeMessages", args)))
}

func wsActionSubscribeMessagesNew(c *websocket.Conn, args string) {
	wsWrite(c, getStrWSJSONTopic(getWSJSONTopic("subscribeMessagesNew", args)))
}

func wsActionUnsubscribeMessagesNew(c *websocket.Conn, args string) {
	wsWrite(c, getStrWSJSONTopic(getWSJSONTopic("unsubscribeMessagesNew", args)))
}

func wsActionSubscribePresences(c *websocket.Conn, args string) {
	wsWrite(c, getStrWSJSONTopic(getWSJSONTopic("subscribePresences", args)))
}

func wsActionUnsubscribePresences(c *websocket.Conn, args string) {
	wsWrite(c, getStrWSJSONTopic(getWSJSONTopic("unsubscribePresences", args)))
}

func wsActionWritePresence(c *websocket.Conn, args string) {
	args = strings.Trim(args, " ")
	j := getWSJSONTopic("writePresence", args[strings.Index(args, " ")+1:])
	s := strings.Split(args, " ")
	j.Status = s[0]
	wsWrite(c, getStrWSJSONTopic(j))
}

func wsActionSubscribeUsers(c *websocket.Conn) {
	wsWrite(c, getStrWSJSONTopic(getWSJSONTopic("subscribeUsers", "")))
}

func wsActionUnsubscribeUsers(c *websocket.Conn) {
	wsWrite(c, getStrWSJSONTopic(getWSJSONTopic("unsubscribeUsers", "")))
}

func wsWrite(c *websocket.Conn, line []byte) {
	if err := c.WriteMessage(1, line); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
