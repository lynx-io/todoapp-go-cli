/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strconv"
)

var (
	id int
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark task as completed",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO:  Find id in task list
		tasks, err := GetAllItems()

		if err != nil {
			fmt.Println("There was a problem getting all items", err)
			return
		}

		homeDir, err := os.UserHomeDir()

		if err != nil {
			fmt.Println("There was a problem getting homedir", err)
			return
		}

		filePath := filepath.Join(homeDir, "lynx-io/databases/tasks.csv")

		file, err := os.Create(filePath)

		if err != nil {
			fmt.Println("There was a problem reading the file", err)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		for _, task := range tasks {
			if task.Id == id {
				task.Completed = true
			}

			fmt.Println(task.Completed)

			err = writer.Write([]string{strconv.Itoa(task.Id), task.Details, strconv.Itoa(task.Urgency), strconv.FormatBool(task.Completed)})
			if err != nil {
				fmt.Println("Error updating task", err)
				return
			}

		}
		fmt.Printf("Task updated\n")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	completeCmd.Flags().IntVar(&id, "id", 0, "Type ID of Task to complete")
}
