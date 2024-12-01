/*
Copyright © 2024 ayrbox <sabin.dangol@hotmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/ayrbox/pass/db"
	"github.com/nrednav/cuid2"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [account_name]",
	Short: "Add account name",
	Long: `Add new account in password manager. 
Account is name of the account that you can remember.

For example, "gmail" is account with username and password`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pm, err := db.Open(dbName)
		if err != nil {
			log.Fatal(err)
		}

		acc := &db.Account{
			Id:   cuid2.Generate(),
			Name: args[0],
		}

		fmt.Printf("%v", *acc)

		_, error := pm.AddAccount(acc)

		if error != nil {
			log.Fatal(error)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
