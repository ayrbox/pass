/*
Copyright © 2024 ayrbox <hello@ayrbox.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var dbName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pass",
	Short: "CLI Password Manager",
	Long: `CLI Password Manager that use sqlite to store.
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dbName, "dbName", "d", "default.db", "Folder path to create sqlite database")
}
