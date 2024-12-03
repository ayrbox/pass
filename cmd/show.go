/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ayrbox/pass/db"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show [accountName]",
	Short: "Show account",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		pm, err := db.Open(dbName)
		if err != nil {
			return fmt.Errorf("Error opening database: %v", err)
		}

		account, err := pm.GetAccountByName(args[0])
		if err != nil {
			return fmt.Errorf("Error getting account: %v", err)
		}

		password, err := pm.GetPassword(&account)
		if err != nil {
			return fmt.Errorf("Error getting passwords: %v", err)
		}

		if account.Username != "" {
			pterm.Printfln("Username: %s", account.Username)
		}
		pterm.Printfln("Password: %s", password)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
