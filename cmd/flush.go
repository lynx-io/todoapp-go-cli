/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// flushCmd represents the flush command
var flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "Removes all items from the Todo list",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			fmt.Println("There was a problem getting homedir", err)
			return
		}

		filePath := filepath.Join(homeDir, "lynx-io/databases/tasks.csv")

		err = os.Remove(filePath)

		if err != nil {
			fmt.Println("There was a problem deleting the file", err)
			return
		}

		fmt.Println("Database flushed, no tasks remaining in the Todo list")
	},
}

func init() {
	rootCmd.AddCommand(flushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
