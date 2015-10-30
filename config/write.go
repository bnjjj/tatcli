package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdConfigTemplate = &cobra.Command{
	Use:   "template",
	Short: "Write a template configuration file in $HOME/.tatcli/config.json: tatcli config template",
	Run: func(cmd *cobra.Command, args []string) {
		writeTemplate()
	},
}

type templateJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
}

func writeTemplate() {
	var templateJSON templateJSON

	if viper.GetString("username") != "" {
		templateJSON.Username = viper.GetString("username")
	}
	if viper.GetString("password") != "" {
		templateJSON.Password = viper.GetString("password")
	}
	if viper.GetString("url") != "" {
		templateJSON.URL = viper.GetString("url")
	}
	jsonStr, err := json.MarshalIndent(templateJSON, "", "  ")
	internal.Check(err)
	jsonStr = append(jsonStr, '\n')
	filename := internal.ConfigFile

	dir := path.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		internal.Check(os.Mkdir(dir, 0740))
	}

	internal.Check(ioutil.WriteFile(filename, jsonStr, 0600))
	fmt.Printf("%s is written\n", filename)
}
