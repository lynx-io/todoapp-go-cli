/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func generateUniqueId() (int, error) {

	tasks, err := GetAllItems()

	if err != nil {
		return 0, err
	}

	var existingIds []int

	existingIds = make([]int, len(tasks))

	for index, task := range tasks {
		existingIds[index] = task.Id
	}

	var randomNumber int

	for {
		randomNumber = rand.IntN(100)
		index := slices.IndexFunc(existingIds, func(id int) bool {
			return id == randomNumber
		})

		if index == -1 {
			break
		}
	}

	return randomNumber, nil
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new item to the list",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		// Prompt for task details
		fmt.Print("Enter task details:\n")
		details, _ := reader.ReadString('\n')
		details = strings.TrimSpace(details)

		// Promt for task urgency (1, 2 or 3)
		fmt.Print("Enter task urgency: ")
		var urgency int
		_, err := fmt.Scan(&urgency)

		if err != nil {
			fmt.Println("Please enter a number from 1 to 3")
			return
		}

		fmt.Printf("Adding task to the list...")

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error opening file:", err)
		}

		filePath := filepath.Join(homeDir, "lynx-io/databases/tasks.csv")

		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println("Error opening file:", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		id, err := generateUniqueId()

		if err != nil {
			fmt.Println("Error generating unique id", err)
			return
		}

		err = writer.Write([]string{strconv.Itoa(id), details, strconv.Itoa(urgency), strconv.FormatBool(false)})

		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Printf("Task added\n")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
