/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/ayrbox/pass/db"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise password manager db.",
	Long:  `Creates password manager database in users home directory`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		pm, err := db.Open(dbName)
		if err != nil {
			log.Fatal(err)
		}
		if err := pm.Init(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("DB: %v has been initialised.\n", dbName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
