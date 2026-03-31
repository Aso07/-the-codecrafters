package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to CLI Calculator")
	fmt.Println("Type 'help' to see commands")
	fmt.Println("Type 'quit' to Exit")
	fmt.Println("Example: add 5 3")

	for {
		var command string

		fmt.Print("> ")
		fmt.Scan(&command)

		parts := strings.Split(command, " ")

		if len(parts) == 0 {
			fmt.Println("Error: Empty command")
			continue
		}

		command = strings.ToLower(command)

		if command == "quit" {
			fmt.Println("Goodbye")
			break
		}
		if command == "help" {
			fmt.Println("add <a> <b>")
			fmt.Println("sub <a> <b>")
			fmt.Println("mul <a> <b>")
			fmt.Println("div <a> <b>")
			fmt.Println("quit")
			continue
		}

		var a, b float64
		_, err := fmt.Scanln(&a, &b)

		if err != nil {
			fmt.Println("Error: Invalid command. Example: add 5 3")
			continue
		}

		switch command {
		case "add":
			fmt.Println("Result:", a+b)

		case "sub":
			fmt.Println("Result:", a-b)

		case "mul":
			fmt.Println("Result:", a*b)

		case "div":
			if b == 0 {
				fmt.Println("Error: Cannot divide by Zero")
				continue
			}
			fmt.Println("Result:", a/b)

		default:
			fmt.Println("Unknown command: Type 'Help'")
		}
	}
}
