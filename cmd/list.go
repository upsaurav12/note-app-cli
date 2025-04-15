/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/upsaurav12/note-app-cli/data"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "See all the list of notes that you have created",
	Long:  `See all the list of notes that you have created`,
	Run: func(cmd *cobra.Command, args []string) {
		data.DisplayAllNotes()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
}
