package cmd

import "github.com/spf13/cobra"

func GetDbName(db *string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		dbName, err := cmd.Flags().GetString("dbName")
		if err != nil {
			*db = "default.db"
		}
		*db = dbName
	}
}
