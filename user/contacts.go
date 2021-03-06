package user

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdUserContacts = &cobra.Command{
	Use:   "contacts",
	Short: "Get contacts presences since n seconds: tatcli user contacts <seconds>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			_, err := strconv.Atoi(args[0])
			if err == nil {
				contacts(args[0])
				return
			}
		}
		fmt.Fprintf(os.Stderr, "Invalid argument: tatcli user contacts --help\n")
	},
}

func contacts(sinceSeconds string) {
	fmt.Print(internal.GetWantReturn("/user/me/contacts/" + sinceSeconds))
}
