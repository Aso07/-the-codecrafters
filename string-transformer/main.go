package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var smallWords = map[string]bool{
	"a": true, "an": true, "the": true, "and": true,
	"but": true, "or": true, "for": true, "nor": true,
	"on": true, "at": true, "to": true, "by": true,
	"in": true, "of": true, "up": true, "as": true,
	"is": true, "it": true,
}

func capitalizeWord(word string) string {
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func upper(text string) string { return strings.ToUpper(text) }
func lower(text string) string { return strings.ToLower(text) }

func cap(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		words[i] = capitalizeWord(w)
	}
	return strings.Join(words, " ")
}

func title(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		w = strings.ToLower(w)
		if i == 0 || !smallWords[w] {
			words[i] = capitalizeWord(w)
		} else {
			words[i] = w
		}
	}
	return strings.Join(words, " ")
}

func snake(text string) string {
	result := ""
	for _, ch := range strings.ToLower(text) {
		if ch == ' ' {
			result += "_"
		} else if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			result += string(ch)
		}
	}
	return result
}

func reverseWord(word string) string {
	r := []rune(word)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverse(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		words[i] = reverseWord(w)
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")
	fmt.Println("──────────────────────────────────────")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		parts := strings.SplitN(input, " ", 2)
		command := strings.ToLower(parts[0])

		if command == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
		}

		if len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
			fmt.Printf("✗ No text provided. Usage: %s <text>\n", command)
			continue
		}

		text := parts[1]

		switch command {
		case "upper":
			fmt.Println("→", upper(text))
		case "lower":
			fmt.Println("→", lower(text))
		case "cap":
			fmt.Println("→", cap(text))
		case "title":
			fmt.Println("→", title(text))
		case "snake":
			fmt.Println("→", snake(text))
		case "reverse":
			fmt.Println("→", reverse(text))
		default:
			fmt.Printf("✗ Unknown command: %q\n", command)
			fmt.Println("Valid commands: upper, lower, cap, title, snake, reverse, exit")
		}
	}
}