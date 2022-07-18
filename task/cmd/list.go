/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long: `List all your todos retrieving them from the database`,
    Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
        tasks, err := db.GetTodos()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("The list of tasks is:")
        for _, task := range tasks {
            fmt.Printf("%v -> %s\n", task.TodoId, task.Name)
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
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
