package config

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdConfigTemplate)
	Cmd.AddCommand(cmdConfigShow)
}

// Cmd config
var Cmd = &cobra.Command{
	Use:     "config",
	Short:   "Config commands: tatcli config --help",
	Long:    `Config commands: tatcli config <command>`,
	Aliases: []string{"c"},
}
