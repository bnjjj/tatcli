package update

import (
	"fmt"
	"net/http"
	"os"

	"github.com/inconshreveable/go-update"
	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

// used by CI to inject architecture (linux-amd64, etc...) at build time
var architecture string

// used by CI to inject url for downloading with tatcli update.
// value of urlUpdate injected at build time
// full URL update is constructed with architecture var :
// urlUpdate + architecture + "/tatcli", tatcli is the binary
var urlUpdateRelease string
var urlUpdateSnapshot string

func init() {
	Cmd.AddCommand(cmdUpdateSnapshot)
}

// Cmd update
var Cmd = &cobra.Command{
	Use:     "update",
	Short:   "Update tatcli to the latest release version: tatcli update",
	Long:    `tatcli update`,
	Aliases: []string{"up"},
	Run: func(cmd *cobra.Command, args []string) {
		doUpdate(fmt.Sprintf("%s%s"+"/tatcli", urlUpdateRelease, architecture))
	},
}

func doUpdate(url string) {
	if internal.Verbose {
		fmt.Printf("Url to update tatcli: %s\n", url)
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error when downloading tatcli: %s\n", err.Error())
		fmt.Printf("Url: %s\n", url)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Error http code: %d, url called: %s\n", resp.StatusCode, url)
		os.Exit(1)
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		fmt.Printf("Error when updating tatcli: %s\n", err.Error())
		fmt.Printf("Url: %s\n", url)
		os.Exit(1)
	}
}
