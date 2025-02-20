package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"go-pg-formatter/internal/format" // TODO - update
)

var rootCmd = &cobra.Command{
	Use:   "pgfmt",
	Short: "Go library for formatting SQL or PostgreSQL files or queries",
	Long:  `This command allows you to format SQL (.sql) and PostgreSQL (.pgsql) files or inline queries using a Go-based formatting library.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path, _ := cmd.Flags().GetString("destination-path")

		// Check if a file was provided
		if path == "" {
			return fmt.Errorf("\033[31mno path specified\033[0m")
		}

		if err := format.Format(path); err != nil {
			return fmt.Errorf("\033[31m%s\033[0m", err.Error())
		}

		return nil
	},
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("destination-path", "d", "", "Destination file path")
}
