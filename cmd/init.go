/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/ayrbox/pass/db"
	"github.com/ayrbox/pass/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise password manager.",
	Long:  `Creates password manager sqlite database in specified location or default location`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := utils.SetupPath("pass_manager")
		if err != nil {
			log.Fatal(err)
		}
		if _, err := db.Init(path, "default.db"); err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	initCmd.Flags().StringP("path", "p", "~", "Folder path to create sqlite database")

	rootCmd.AddCommand(initCmd)
}
