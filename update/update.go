package update

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"github.com/inconshreveable/go-update"
	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

// used by CI to inject architecture (linux-amd64, etc...) at build time
var architecture string

func init() {
	if urlUpdateSnapshot != "" {
		Cmd.AddCommand(cmdUpdateSnapshot)
	}
}

// Cmd update
var Cmd = &cobra.Command{
	Use:     "update",
	Short:   "Update tatcli to the latest release version: tatcli update",
	Long:    `tatcli update`,
	Aliases: []string{"up"},
	Run: func(cmd *cobra.Command, args []string) {
		doUpdate("", architecture)
	},
}

func getURLArtifactFromGithub(architecture string) string {
	client := github.NewClient(nil)
	release, resp, err := client.Repositories.GetLatestRelease("ovh", "tatcli")
	if err != nil {
		fmt.Printf("Repositories.GetLatestRelease returned error: %v\n%v", err, resp.Body)
		os.Exit(1)
	}

	if len(release.Assets) > 0 {
		for _, asset := range release.Assets {
			if *asset.Name == "tatcli-"+architecture {
				return *asset.BrowserDownloadURL
			}
		}
	}

	fmt.Println("Invalid Artifacts on latest release. Please try in few minutes.")
	fmt.Println("If it's persit, please open an issue on https://github.com/ovh/tatcli/issues")
	os.Exit(1)
	return ""
}

func doUpdate(baseurl, architecture string) {
	if architecture == "" {
		fmt.Println("You seem to have a custom build of tatcli")
		fmt.Println("Please download latest release on https://github.com/ovh/tatcli/releases")
		os.Exit(1)
	}
	url := getURLArtifactFromGithub(architecture)
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

	fmt.Printf("Getting latest release from : %s ...\n", url)
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		fmt.Printf("Error when updating tatcli: %s\n", err.Error())
		fmt.Printf("Url: %s\n", url)
		os.Exit(1)
	}
	fmt.Println("Updating done.")
}
