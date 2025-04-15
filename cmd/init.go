/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/upsaurav12/note-app-cli/data"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "inti command for notes app",
	Long:  `inti command for notes app`,
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
