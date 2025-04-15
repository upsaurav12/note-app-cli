/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/upsaurav12/note-app-cli/cmd"
	"github.com/upsaurav12/note-app-cli/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
