package main

import (
	"choose-your-own-adventure/cli"
	"choose-your-own-adventure/website"
	"fmt"
	"log"
)

func main() {
    // Show options
    fmt.Println("Welcome to Choose your own adveture!")
    fmt.Println("0. Cli mode")
    fmt.Println("1. Web mode")
    fmt.Println("Choose your option: ")

    // Ask option
    var opt int
    _, err := fmt.Scanln(&opt)
    if err != nil {
        log.Fatal("Not a valid integer", err)
    }

    // Run option
    if opt == 0 {
        cli.Start()
    } else if opt == 1 {
        log.Fatal(website.Start())
    } else {
        log.Fatal("Not a valid option")
    }
}
