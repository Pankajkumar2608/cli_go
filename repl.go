package main 

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()
		
		cleaned := cleanInput(text);
		if len(cleaned) == 0 {
			continue
		}
		
		commandName := cleaned[0]
		
		availableCommand := getCommand()
		
		command, ok := availableCommand[commandName]
		
		if !ok {
			fmt.Println("Incorrect Command bitch")
            continue
		}
		
		command.callback();
		
	}
	
}

type cliCommand struct {
	name   string
	description  string
	callback   func() error
}

func getCommand () map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name : "help",
			description: "Prints the help menu",
			callback: callbackHelp ,
		},
		"exit": {
			name: "exit",
			description:"Exit the program",
			callback: callbackExit,
		},
	}
}

func cleanInput (str string) []string {
	lowered := strings.ToLower(str);
	words := strings.Fields(lowered)
	return words;
}