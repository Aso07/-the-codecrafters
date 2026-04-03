// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Agene Okoh]
// Squad:  Goroutines

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: Goroutines
// ───────────────────────────────────────────
//
// Input line types:
//
//   1. Normal report lines
//   2. Lines in ALL CAPS
//   3. Lines in all lowercase
//   4. Lines starting with TODO:
//   5. Lines starting with CLASSIFIED:
//   6. Lines that are only dashes or blank
//   7. Lines with leading/trailing whitespace
//   8. Lines containing numbers and symbols
//
// Transformation rules (in order):
//
//   1. Trim all leading and trailing whitespace
//   2. Remove lines that are only dashes or blank
//   3. Replace TODO: with ✦ ACTION:
//   4. Replace CLASSIFIED: with [REDACTED]:
//   5. Reverse the words in any line that contains the word REVERSE
//
// Output format:
//
//   Header: yes — SENTINEL FIELD REPORT — PROCESSED
//   Line numbering format: 001., 002., 003. (three-digit zero-padded)
//   Summary block: yes — Lines Processed, Lines Written, Lines Removed
//
// Terminal summary fields:
//
//   ✦ Lines read    : <number>
//   ✦ Lines written : <number>
//   ✦ Lines removed : <number>
//   ✦ Rules applied : Trim whitespace, Remove blank/dash lines, Replace TODO, Replace CLASSIFIED, Reverse REVERSE lines
// ═══════════════════════════════════════════

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
		fmt.Println(" Input and output cannot be the same file.")
		return
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf(" File not found: %s\n", inputFile)
		return
	}
	defer inFile.Close()

	if info, err := os.Stat(outputFile); err == nil && info.IsDir() {
		fmt.Println(" Cannot write to output: path is a directory, not a file.")
		return
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println(" Cannot write to output file.")
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

		if shouldRemove(line) {
			linesRemoved++
			continue
		}

		line = replaceTODO(line)
		line = replaceClassified(line)
		line = reverseIfREVERSE(line)

		linesWritten++
		fmt.Fprintln(writer, addLineNumber(line, linesWritten))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(" Error reading file.")
		return
	}

	if linesRead == 0 {
		fmt.Println(" Input file is empty. Nothing to process.")
	}

	fmt.Println("* Lines read    :", linesRead)
	fmt.Println("* Lines written :", linesWritten)
	fmt.Println("* Lines removed :", linesRemoved)
	fmt.Println("* Rules applied : Trim whitespace, Remove blank/dash lines, Replace TODO, Replace CLASSIFIED, Reverse REVERSE lines")
}

func trim(line string) string {
	return strings.TrimSpace(line)
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

func replaceTODO(line string) string {
	return strings.ReplaceAll(line, "TODO:", " ACTION:")
}

func replaceClassified(line string) string {
	return strings.ReplaceAll(line, "CLASSIFIED:", "[REDACTED]:")
}

func reverseIfREVERSE(line string) string {
	if !strings.Contains(line, "REVERSE") {
		return line
	}
	words := strings.Fields(line)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func addLineNumber(line string, index int) string {
	return fmt.Sprintf("%03d. %s", index, line)
}
