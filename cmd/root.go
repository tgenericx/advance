/*
Copyright © 2026 tgenericx <hammedanuoluwapopelumi@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "advance",
	Short: "A minimal, local-first CLI todo manager.",
	Long: `Advance is a fast, minimal CLI tool for managing todos from your terminal.

It is designed to help you move work forward with clarity and momentum.`,
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
}
