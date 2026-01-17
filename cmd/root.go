/*
Copyright © 2026 tgenericx <hammedanuoluwapopelumi@gmail.com>
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tgenericx/advance/todo"
)

var rootCmd = &cobra.Command{
	Use:   "advance",
	Short: "A minimal, local-first CLI todo manager.",
	Long: `Advance is a fast, minimal CLI tool for managing todos from your terminal.

It is designed to help you move work forward with clarity and momentum.`,
}

var (
	datafile string
	store    todo.Store
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		cobra.CheckErr(err)
	}

	rootCmd.PersistentFlags().StringVar(
		&datafile,
		"datafile",
		filepath.Join(home, ".advance.json"),
		"path to the data file",
	)

	cobra.OnInitialize(initStore)
}

func initStore() {
	store = todo.NewFileStore(datafile)
}
