/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"task/cmd"
	"task/db"
)

func main() {
    db.Open()
    defer db.Close()
	cmd.Execute()
}
