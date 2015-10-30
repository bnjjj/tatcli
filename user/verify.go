package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var save bool

func init() {
	cmdUserVerify.Flags().BoolVarP(&save, "save", "s", false, "Save configuration after verify in $HOME/.tatcli/config.json")
}

var cmdUserVerify = &cobra.Command{
	Use:   "verify",
	Short: "Verify account: tatcli user verify [--save] <username> <tokenVerify>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			userVerify(args[0], args[1])
		} else {
			fmt.Println("Invalid argument to verify account: tatcli user verify --help")
		}
	},
}

type verifyJSON struct {
	Message  string `json:"message,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	URL      string `json:"url,omitempty"`
}

func userVerify(username, tokenVerify string) {
	url := fmt.Sprintf("/user/verify/%s/%s", username, tokenVerify)
	r := internal.GetWantReturn(url)

	var verifyJSON verifyJSON
	err := json.Unmarshal([]byte(r), &verifyJSON)
	internal.Check(err)

	// Display return to user (contains clear password)
	fmt.Printf(r)

	if save && verifyJSON.Password != "" && verifyJSON.Username != "" && verifyJSON.URL != "" {
		verifyJSON.Message = ""
		jsonStr, err := json.MarshalIndent(verifyJSON, "", "  ")
		internal.Check(err)
		jsonStr = append(jsonStr, '\n')
		filename := internal.ConfigFile
		dir := path.Dir(filename)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			internal.Check(os.Mkdir(dir, 0740))
		}
		internal.Check(ioutil.WriteFile(filename, jsonStr, 0600))
	}
}
