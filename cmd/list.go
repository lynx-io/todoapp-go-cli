/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	priority bool
)

type Task struct {
	Id      int
	Details string
	Urgency int
}

func sortTaskByUrgency(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Urgency < tasks[j].Urgency
	})
}

func GetAllItems() ([]Task, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println("There was a problem getting homedir", err)
		return nil, err
	}

	filePath := filepath.Join(homeDir, "lynx-io/databases/tasks.csv")

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("There was a problem reading the file", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var tasks []Task

	for _, record := range records {
		id, err := strconv.Atoi(record[0])

		if err != nil {
			return nil, err
		}

		urgency, err := strconv.Atoi(record[2])

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, Task{
			Id:      id,
			Details: record[1],
			Urgency: urgency,
		})
	}

	return tasks, nil
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks of the todo list",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := GetAllItems()

		if err != nil {
			fmt.Println("There was a problem getting all items", err)
			return
		}

		if priority {
			fmt.Println("Priority set")
			sortTaskByUrgency(tasks)
		}

		fmt.Println("### Todo list ###")

		for _, task := range tasks {
			fmt.Printf("Urgency : %d | Task: %s\n", task.Urgency, task.Details)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().BoolVarP(&priority, "sort", "s", false, "If set, then sort by urgency")
}
