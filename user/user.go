package user

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cmdUserList)
	Cmd.AddCommand(cmdUserMe)
	Cmd.AddCommand(cmdUserContacts)
	Cmd.AddCommand(cmdUserAddContact)
	Cmd.AddCommand(cmdUserRemoveContact)
	Cmd.AddCommand(cmdUserAddFavoriteTopic)
	Cmd.AddCommand(cmdUserRemoveFavoriteTopic)
	Cmd.AddCommand(cmdUserEnableNotificationsTopic)
	Cmd.AddCommand(cmdUserDisableNotificationsTopic)
	Cmd.AddCommand(cmdUserAddFavoriteTag)
	Cmd.AddCommand(cmdUserRemoveFavoriteTag)
	Cmd.AddCommand(cmdUserAdd)
	Cmd.AddCommand(cmdUserReset)
	Cmd.AddCommand(cmdUserResetSystem)
	Cmd.AddCommand(cmdUserConvertToSystem)
	Cmd.AddCommand(cmdUserArchive)
	Cmd.AddCommand(cmdUserRename)
	Cmd.AddCommand(cmdUserUpdate)
	Cmd.AddCommand(cmdUserSetAdmin)
	Cmd.AddCommand(cmdUserVerify)
}

// Cmd user
var Cmd = &cobra.Command{
	Use:     "user",
	Short:   "User commands: tatcli user --help",
	Long:    `User commands: tatcli user <command>`,
	Aliases: []string{"u"},
}

func userPUT(path string) {
	fmt.Print(internal.PutWantReturn("/user" + path))
}

func userPOST(path string) {
	internal.PostWant("/user"+path, nil)
}

func userDelete(path string) {
	internal.DeleteWant("/user"+path, nil)
}
