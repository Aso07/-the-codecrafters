package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Aso Base Converter")
	fmt.Println("Use: convert <value> <base>")
	fmt.Println("Type 'quit' to exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			fmt.Println("Goodbye")
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		if strings.ToLower(input) == "quit" {
			fmt.Println("Goodbye")
			break
		}

		parts := strings.Split(input, " ")

		if len(parts) != 3 {
			fmt.Println("Usage: convert <value> <base>")
			continue
		}

		command := strings.ToLower(parts[0])
		value := parts[1]
		base := strings.ToLower(parts[2])

		if command != "convert" {
			fmt.Println("Usage: convert <value> <base>")
			continue
		}

		switch base {

		case "bin":
			num, err := strconv.ParseInt(value, 2, 64)
			if err != nil {
				fmt.Println("Error: invalid binary number")
				continue
			}
			fmt.Println(" Decimal:", num)

		case "hex":
			num, err := strconv.ParseInt(value, 16, 64)
			if err != nil {
				fmt.Println("Error: invalid hex number")
				continue
			}
			fmt.Println(" Decimal:", num)

		case "dec":
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error: invalid decimal number")
				continue
			}

			binary := strconv.FormatInt(int64(num), 2)
			hex := strings.ToUpper(strconv.FormatInt(int64(num), 16))

			fmt.Println(" Binary:", binary)
			fmt.Println(" Hex:   ", hex)

		default:
			fmt.Println("Error: base must be bin, hex, or dec")
		}
	}
}
