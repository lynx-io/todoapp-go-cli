/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	completedOnly bool
)

// flushCmd represents the flush command
var flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "Removes all items from the Todo list",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := GetAllItems()

		if err != nil {
			fmt.Println("There was an error getting all items", err)
			return
		}

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

		if completedOnly {
			writer, err := os.Create(filePath)
			defer writer.Close()

			if err != nil {
				fmt.Println("Error creating file", err)
				return
			}

			csv := csv.NewWriter(writer)
			defer csv.Flush()

			for _, task := range tasks {
				if !task.Completed {
					err = csv.Write([]string{strconv.Itoa(task.Id), task.Details, strconv.Itoa(task.Urgency), strconv.FormatBool(task.Completed)})

					if err != nil {
						fmt.Println("Error writing", err)
						return
					}
				}
			}
			fmt.Println("Completed only items removed")
		} else {

			fmt.Println("Database flushed, no tasks remaining in the Todo list")
		}

	},
}

func init() {
	rootCmd.AddCommand(flushCmd)

	flushCmd.Flags().BoolVarP(&completedOnly, "completed", "c", false, "Help message for toggle")
}
