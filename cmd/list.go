/*
Copyright © 2026 tgenericx <hammedanuoluwapopelumi@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo items",
	Run:   runList,
}

func runList(cmd *cobra.Command, args []string) {
	items, err := store.Load()
	if err != nil {
		cobra.CheckErr(err)
	}

	for _, item := range items {
		status := " "
		if item.Done {
			status = "x"
		}

		fmt.Printf("%d [%s] %s\n", item.ID, status, item.Text)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
