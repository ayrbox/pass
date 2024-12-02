/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ayrbox/pass/db"
	"github.com/spf13/cobra"
)

var (
	accountName string
	username    string
	password    string
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [accountName]",
	Short: "Update given account",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		pm, err := db.Open(dbName)
		if err != nil {
			return err
		}

		account, _ := pm.GetAccountByName(args[0])

		// TODO: confirm before actually updating
		if accountName != "" {
			pm.UpdateAccountName(&account, accountName)
		}

		if username != "" {
			pm.UpdateUsername(&account, username)
		}

		if password != "" {
			pm.UpdatePassword(&account, password)
		}

		fmt.Println("Updated")
		return nil
	},
}

func init() {
	updateCmd.Flags().StringVarP(&accountName, "accountName", "a", "", "New account name")
	updateCmd.Flags().StringVarP(&username, "username", "u", "", "New username")
	updateCmd.Flags().StringVarP(&password, "password", "p", "", "New password")
	rootCmd.AddCommand(updateCmd)
}
