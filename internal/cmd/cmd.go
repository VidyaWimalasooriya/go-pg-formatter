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
		arguments := getConfig(cmd.Flags())

		// Check if a file was provided
		if path == "" {
			return fmt.Errorf("\033[31mno path specified\033[0m")
		}

		if err := format.Format(path, arguments...); err != nil {
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
	rootCmd.Flags().StringP("destination-path", "d", "", "Destination file path")
	rootCmd.Flags().BoolP("anonymize", "a", false, "obscure all literals in queries, useful to hide confidential data before formatting.")
	rootCmd.Flags().BoolP("comma-start", "b", false, "in a parameters list, start with the comma.")
	rootCmd.Flags().BoolP("comma-end", "e", false, "in a parameters list, end with the comma.")
	rootCmd.Flags().IntP("keyword-case", "u", 2, "Change the case of the reserved keyword. Default is uppercase: 2. Values: 0=>unchanged, 1=>lowercase, 2=>uppercase, 3=>capitalize.")
	rootCmd.Flags().BoolP("no-rcfile", "X", false, "don't read rc files automatically (./.pg_format or $HOME/.pg_format or $XDG_CONFIG_HOME/pg_format).")
	rootCmd.Flags().IntP("spaces", "s", 4, "change space indent, default 4 spaces.")
	rootCmd.Flags().BoolP("nocomment", "n", false, "remove any comment from SQL code.")
	rootCmd.Flags().BoolP("tabs", "T", false, "use tabs instead of space characters, when used spaces is set to 1 whatever is the value set to -s.")
	rootCmd.Flags().BoolP("comma-break", "B", false, "in insert statement, add a newline after each comma.")

	//TODO - Support for more flags
}
