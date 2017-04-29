// Copyright © 2017 Christian H Hercules <cxhercules@gmail.com>
//

package cmd

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export <profile>",
	Short: "Export the aws credentials to be read in by environment",
	Long: `Export aws credentials to be set as environment variables
If no argument is passed the default profile is exported.
`,
	Example: `
aws-env export testenv
export AWS_ACCESS_KEY_ID=XXXXXXXXXXXXXXXX
export AWS_SECRET_ACCESS_KEY=YYYYYYYYYYYYYYYYYYYYYYYYYYY
`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		var profileName string
		cfg, err := ini.Load(credentialsFilename())
		check(err)

		if len(args) == 0 {
			profileName = "default"
		} else if len(args) == 1 {
			profileName = args[0]
		} else {
			fmt.Println("Too many arguments passed, failing ...")
			os.Exit(-1)
		}

		profile, err := cfg.GetSection(profileName)
		check(err)

		key, err := profile.GetKey("aws_access_key_id")
		fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", key)
		check(err)

		secret, err := profile.GetKey("aws_secret_access_key")
		fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", secret)
		check(err)
	},
}

func init() {
	RootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
