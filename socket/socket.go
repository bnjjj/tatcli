package socket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ovh/tat/models"
	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Cmd.AddCommand(cmdSocketDump)
	Cmd.AddCommand(cmdSocketMessages)
	Cmd.AddCommand(cmdSocketMessagesNew)
	Cmd.AddCommand(cmdSocketInteractive)
	Cmd.AddCommand(cmdSocketUsers)
}

// Cmd socket
var Cmd = &cobra.Command{
	Use:     "socket",
	Short:   "Socket commands: tatcli socket --help",
	Long:    `Socket commands: tatcli socket [<command>]`,
	Aliases: []string{"s"},
}

func wsActionConnect(c *websocket.Conn) {
	w := &models.WSConnectJSON{}
	w.Username = viper.GetString("username")
	w.Password = viper.GetString("password")
	jsonStr, err := json.Marshal(w)
	if err != nil {
		fmt.Printf("Error while converting to json (connect): %s", err.Error())
	}
	wsWrite(c, jsonStr)
}
func newClient() *websocket.Conn {
	internal.ReadConfig()
	if viper.GetString("url") == "" {
		fmt.Println("Invalid Configuration: invalid URL. See tatcli config --help")
		os.Exit(1)
	}

	i := strings.Index(viper.GetString("url"), ":")
	endURL := viper.GetString("url")[i:] + "/socket/ws"

	url := "ws" + endURL
	dialer := websocket.DefaultDialer
	if internal.IsHTTPS() {
		url = "wss" + endURL
		dialer = &websocket.Dialer{
			TLSClientConfig:  internal.GetTLSConfig(),
			HandshakeTimeout: time.Duration(time.Duration(time.Second) * 9),
		}
	}

	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Add("Content-Type", "application/json")

	c, _, err := dialer.Dial(url, r.Header)
	internal.Check(err)
	fmt.Printf("Connected to %s [Ctrl+c to quit]\n", url)

	wsActionConnect(c)
	return c
}
