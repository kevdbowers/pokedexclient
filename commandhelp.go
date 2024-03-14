package main

import "fmt"

func commandHelp(config *configuration) error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("\nCommands:\n")
	for _, c := range checkInput() {
		fmt.Printf("	%s:  %s\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}
