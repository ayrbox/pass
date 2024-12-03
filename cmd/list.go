/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ayrbox/pass/db"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the accounts",
	RunE: func(cmd *cobra.Command, args []string) error {
		pm, err := db.Open(dbName)
		if err != nil {
			return err
		}

		accounts, err := pm.GetAccounts()
		accountData := pterm.TableData{
			{"ID", "Name", "Username"},
		}

		for _, account := range accounts {
			accountData = append(accountData, []string{account.Id[:10], account.Name, account.Username})
		}

		pterm.DefaultTable.WithHasHeader().WithData(accountData).Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
