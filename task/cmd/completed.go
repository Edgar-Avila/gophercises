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

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Show a list with the completed tasks today",
    Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
        tasks, err := db.GetCompletedToday()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("The list of completed tasks is:")
        for _, task := range tasks {
            fmt.Printf("%v -> %s\n", task.TodoId, task.Name)
        }
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
