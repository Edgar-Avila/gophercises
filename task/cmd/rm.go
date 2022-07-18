/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Use this command to completely delete a task",
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        id, err := strconv.ParseUint(args[0], 10, 64)
        if err != nil {
            log.Fatal(err)
        }
        if err = db.RemoveTodo(id); err != nil {
            log.Fatal(err)
        }
        fmt.Println("Task removed successfully")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
