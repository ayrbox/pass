/*
Copyright © 2024 ayrbox <sabin.dangol@hotmail.com>
*/
package cmd

import (
	"errors"
	"log"

	"github.com/ayrbox/pass/db"
	"github.com/nrednav/cuid2"
	"github.com/spf13/cobra"
)

func parseArgs(args []string) (string, string) {

	if len(args) <= 0 {
		errors.New("At least one argument is required for account name")
	}

	if len(args) == 1 {
		return args[0], ""
	}

	return args[0], args[1]

}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [account_name] [username]",
	Short: "Add account name",
	Long: `Add new account in password manager. 
Account is name of the account that you can remember.

For example, "gmail" is account with username and password`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pm, err := db.Open(dbName)
		if err != nil {
			log.Fatal(err)
		}

		name, username := parseArgs(args)
		acc := &db.Account{
			Id:       cuid2.Generate(),
			Name:     name,
			Username: username,
		}
		_, error := pm.AddAccount(acc)

		if error != nil {
			log.Printf("Error adding account.\n")
			log.Fatal(error)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
