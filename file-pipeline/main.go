/ ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: Gophers
// ───────────────────────────────────────────
// Input line types:
// Number of lines: 20
// Normal report lines
// Lines in ALL CAPS
// Lines in all lowercase
// Lines starting with TODO:
// Lines with extra leading/trailing spaces

// Transformation rules (in order):
// 1. Trim all leading and trailing whitespace
// 2. Replace TODO: with ✦ ACTION:
// 3. Convert ALL CAPS lines to Title Case
// 4. Convert all lowercase lines to uppercase
// 5. Reverse the words in any line that contains the word REVERSE

// Output format:
// Header: Yes, Exact Text: "Gopher's Sentinel Field Report - Processed"
// Line numbering format : "1."
// Summary block: yes
//
//		Fields :
//			✦ Lines read    :
//			✦ Lines written :
//			✦ Lines removed :
//			✦ Rules applied : [our 5 rules]
//
// Terminal summary fields:
//
//	✦ Lines read    :
//	✦ Lines written :
//	✦ Lines removed :
//	✦ Rules applied : [our 5 rules]
//
// ═══════════════════════════════════════════

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TrimWhitespace(line string) string {
	return strings.TrimSpace(line)
}

func ReplaceTODO(line string) string {
	return strings.ReplaceAll(line, "TODO:", "ACTION:")
}

func AllCapsToTitle(line string) string {
	if line == strings.ToUpper(line) && len(line) > 0 {
		words := strings.Fields(strings.ToLower(line))
		for i := range words {
			words[i] = strings.Title(words[i])
		}
		return strings.Join(words, " ")
	}
	return line
}

func Lower(line string) string {
	return strings.ToLower(line)
}

func ReverseWordsIfContainsReverse(line string) string {
	if !strings.Contains(strings.ToLower(line), "reverse") {
		return line
	}
	words := strings.Fields(line)
	for i := range words {
		runes := []rune(words[i])
		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
			runes[j], runes[k] = runes[k], runes[j]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	in, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("File not found: %s\n", inputFile)
		return
	}
	defer in.Close()

	scanner := bufio.NewScanner(in)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var processed []string
	transformations := []func(string) string{
		TrimWhitespace,
		ReplaceTODO,
		AllCapsToTitle,
		Lower,
		ReverseWordsIfContainsReverse,
	}

	for _, line := range lines {
		for _, fn := range transformations {
			line = fn(line)
		}
		processed = append(processed, line)
	}

	for i := range processed {
		processed[i] = fmt.Sprintf("%d. %s", i+1, processed[i])
	}

	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Cannot write to output: %s\n", outputFile)
		return
	}
	defer out.Close()

	fmt.Fprintln(out, "Gopher's Sentinel Field Report - Processed")
	for _, line := range processed {
		fmt.Fprintln(out, line)
	}

	linesRead := len(lines)
	linesRemoved := 0

	fmt.Fprintln(out, "\n-----summary-----")
	fmt.Fprintf(out, "Lines read    : %d\n", linesRead)
	fmt.Fprintf(out, "Lines written : %d\n", len(processed))
	fmt.Fprintf(out, "Lines removed : %d\n", linesRemoved)
	fmt.Fprintf(out, "Rules applied : TrimWhitespace, ReplaceTODO, AllCapsToTitle, Lower, reverseWord\n")

	fmt.Println("\n--- Terminal Summary ---")
	fmt.Printf("Lines read    : %d\n", linesRead)
	fmt.Printf("Lines written : %d\n", len(processed))
	fmt.Printf("Lines removed : %d\n", linesRemoved)
	fmt.Printf("Rules applied : TrimWhitespace, ReplaceTODO, AllCapsToTitle, lower, reverseWord\n")
}