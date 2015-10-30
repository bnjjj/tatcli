package socket

import "github.com/spf13/cobra"

var cmdSocketUsers = &cobra.Command{
	Use:     "users",
	Short:   "Open websocket and get events users (admin only): tatcli socket users",
	Aliases: []string{"u"},
	Run: func(cmd *cobra.Command, args []string) {
		socketUsers()
	},
}

func socketUsers() {
	c := newClient()
	wsActionSubscribeUsers(c)
	done := make(chan bool)
	go socketRead(c)
	<-done
}
