package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Println("✗ Input and output cannot be the same file.")
		return
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("✗ File not found: %s\n", inputFile)
		return
	}
	defer inFile.Close()

	if info, err := os.Stat(outputFile); err == nil && info.IsDir() {
		fmt.Println("✗ Cannot write to output: path is a directory, not a file.")
		return
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("✗ Cannot write to output file.")
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	fmt.Fprintln(writer, "SENTINEL FIELD REPORT — PROCESSED")

	linesRead := 0
	linesWritten := 0
	linesRemoved := 0

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		linesRead++

		line = trim(line)
		line = replaceTODO(line)
		line = replaceClassified(line)

		if shouldRemove(line) {
			linesRemoved++
			continue
		}

		linesWritten++
		line = addLineNumber(line, linesWritten)

		fmt.Fprintln(writer, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("✗ Error reading file.")
		return
	}

	if linesRead == 0 {
		fmt.Println("⚠ Input file is empty. Nothing to process.")
	}

	fmt.Println("✦ Lines read    :", linesRead)
	fmt.Println("✦ Lines written :", linesWritten)
	fmt.Println("✦ Lines removed :", linesRemoved)
	fmt.Println("✦ Rules applied : trim, replaceTODO, replaceClassified, removeEmpty, addLineNumbers")
}

func trim(line string) string {
	return strings.TrimSpace(line)
}

func replaceTODO(line string) string {
	return strings.ReplaceAll(line, "TODO:", "✦ ACTION:")
}

func replaceClassified(line string) string {
	return strings.ReplaceAll(line, "CLASSIFIED:", "[REDACTED]:")
}

func shouldRemove(line string) bool {
	if line == "" {
		return true
	}

	for _, ch := range line {
		if ch != '-' {
			return false
		}
	}
	return true
}

func addLineNumber(line string, index int) string {
	return fmt.Sprintf("%1d. %s", index, line)
}