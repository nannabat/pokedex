package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliName string = "pokedex"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func printPrompt() {
	fmt.Print(cliName, "> ")
}

func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}

func commandHelp() error {
	commands := availableCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for command := range commands {
		fmt.Printf("%v:%v\n", commands[command].name, commands[command].description)
	}
	return nil

}

func commandExit() error {
	os.Exit(0)
	return nil

}

func availableCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		commands := availableCommands()
		if command, exists := commands[text]; exists {
			command.callback()
		} else {
			commandHelp()
		}
		println()
		printPrompt()
	}
	fmt.Println()

}
