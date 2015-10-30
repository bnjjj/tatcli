package main

import (
	"os"

	"github.com/ovh/tatcli/config"
	"github.com/ovh/tatcli/group"
	"github.com/ovh/tatcli/internal"
	"github.com/ovh/tatcli/message"
	"github.com/ovh/tatcli/presence"
	"github.com/ovh/tatcli/socket"
	"github.com/ovh/tatcli/stats"
	"github.com/ovh/tatcli/topic"
	"github.com/ovh/tatcli/update"
	"github.com/ovh/tatcli/user"
	"github.com/ovh/tatcli/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var home = os.Getenv("HOME")

var rootCmd = &cobra.Command{
	Use:   "tatcli",
	Short: "Text And Tags - Command Line Tool",
	Long:  `Text And Tags - Command Line Tool`,
}

func main() {
	addCommands()
	rootCmd.PersistentFlags().BoolVarP(&internal.Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&internal.Pretty, "pretty", "t", false, "Pretty Print Json Output")
	rootCmd.PersistentFlags().BoolVarP(&internal.SSLInsecureSkipVerify, "sslInsecureSkipVerify", "l", false, "Skip certificate check with SSL connection")
	rootCmd.PersistentFlags().StringVarP(&internal.URL, "url", "", "", "URL Tat Engine, facultative if you have a "+home+"/.tatcli/config.json file")
	rootCmd.PersistentFlags().StringVarP(&internal.Username, "username", "u", "", "username, facultative if you have a "+home+"/.tatcli/config.json file")
	rootCmd.PersistentFlags().StringVarP(&internal.Password, "password", "p", "", "password, facultative if you have a "+home+"/.tatcli/config.json file")
	rootCmd.PersistentFlags().StringVarP(&internal.ConfigFile, "configFile", "c", home+"/.tatcli/config.json", "configuration file, default is "+home+"/.tatcli/config.json")

	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))

	rootCmd.Execute()
}

//AddCommands adds child commands to the root command rootCmd.
func addCommands() {
	rootCmd.AddCommand(config.Cmd)
	rootCmd.AddCommand(group.Cmd)
	rootCmd.AddCommand(message.Cmd)
	rootCmd.AddCommand(presence.Cmd)
	rootCmd.AddCommand(socket.Cmd)
	rootCmd.AddCommand(stats.Cmd)
	rootCmd.AddCommand(topic.Cmd)
	rootCmd.AddCommand(update.Cmd)
	rootCmd.AddCommand(user.Cmd)
	rootCmd.AddCommand(version.Cmd)
}
