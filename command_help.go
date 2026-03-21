package main

import "fmt"

func callbackHelp () error {
	fmt.Println("Welcome to my program bitch")
	fmt.Println("Here are your available commands")
	availableCommands := getCommand()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	
	return nil
}