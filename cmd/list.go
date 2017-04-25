// Copyright © 2017 Christian H Hercules <cxhercules@gmail.com>
//

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
	"github.com/spf13/cobra"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func credentialsFilename() string {
	return filepath.Join(userHomeDir(), ".aws", "credentials")
}

func configFilename() string {
	return filepath.Join(userHomeDir(), ".aws", "config")
}

func userHomeDir() string {
	homeDir := os.Getenv("HOME") // *nix
	if len(homeDir) == 0 {       // windows
		homeDir = os.Getenv("USERPROFILE")
	}

	return homeDir
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the aws profiles configured",
	Long: `List all available aws profiles so that we can use one of those to export:

aws-env list
prod
dev
stage
something

aws-env export stage`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		cfg, err := ini.Load(credentialsFilename())
		check(err)

		// sections := cfg.Sections()
		profiles := cfg.SectionStrings()

		// fmt.Println(sections)
		// fmt.Printf("%v", profiles)

		for _, profile := range profiles {
			if profile != "DEFAULT" {
				fmt.Printf("%s\n", profile)
			}
		}

		fmt.Println("list called")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
