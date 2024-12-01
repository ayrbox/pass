/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ayrbox/pass/db"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen [accountName]",
	Short: "Generate password",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		pm, err := db.Open(dbName)
		if err != nil {
			return err
		}

		account, _ := pm.GetAccountByName(args[0])

		err = pm.GeneratePassword(&account)

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
