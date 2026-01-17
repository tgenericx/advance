/*
Copyright © 2026 tgenericx <hammedanuoluwapopelumi@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tgenericx/advance/todo"
)

var addCmd = &cobra.Command{
	Use:   "add <text>",
	Short: "Add a new todo item",
	Args:  cobra.MinimumNArgs(1),
	Run:   runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	items, err := store.Load()
	if err != nil {
		cobra.CheckErr(err)
	}

	nextID := len(items) + 1

	for _, text := range args {
		items = append(items, todo.Item{
			ID:   nextID,
			Text: text,
			Done: false,
		})
		nextID++
	}

	if err := store.Save(items); err != nil {
		cobra.CheckErr(err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
