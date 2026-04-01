package main

import(
	"fmt"
	"strings"
	"unicode"
	"bufio"
	"os"

)

var smallWords = []string{
	"a", "an", "the", "and", "but", "or", "for", "nor",
	"on", "at", "to", "by", "in", "of", "up", "as",
	"is", "yet", "it",
}

func checkSmallWords(word string) bool {
	for _, w := range smallWords {
		if w == word {
			return true
		}
	}
	return false
}

func title(word string) string {
	words := strings.Fields(word)

	for i, w := range words {
		lower := strings.ToLower(w)

		if i == 0 || !checkSmallWords(lower) {
			runes := []rune(lower)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		} else {
			words[i] = lower
		}
	}

	return strings.Join(words, " ")
}

func upper(word string) string {
	return strings.ToUpper(word)
}

func lower(word string) string {
	return strings.ToLower(word)
}

func cap(word string) string {
	words := strings.Fields(word)

	for i, ch := range words {
		runes := []rune(strings.ToLower(ch))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)

	}
	return strings.Join(words, " ")
}
func snake(word string) string {
	var result strings.Builder

	for _, ch := range word {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || unicode.IsSpace(ch) {
			if unicode.IsLetter(ch) {
				result.WriteRune(unicode.ToLower(ch))
			} else {
				result.WriteRune(ch)
			}
		}
	}

	word = result.String()
	word = strings.Join(strings.Fields(word), "_")
	return word
}

func reverse(word string) string {
	words := strings.Fields(word)
	for i, word := range words {
		runes := []rune(word)
		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}
		words[i] = string(runes)

	}
	return strings.Join(words, " ")
}
func stringTransformer() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")

		fmt.Print(">> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			fmt.Println("Please enter a valid command -- Usage: upper <text>")
			continue
		}

		textInput := strings.Fields(line)
		command := strings.ToLower(textInput[0])

		if command == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			return
		}

		if len(textInput) < 2 {
			fmt.Printf("No text provided. Usage: %s text\n\n", command)
			continue
		}

		text := strings.Join(textInput[1:], " ")

		switch command {
		case "upper":
			fmt.Println("→ ", upper(text))

		case "snake":
			fmt.Println("→ ", snake(text))

		case "lower":
			fmt.Println("→ ", lower(text))

		case "reverse":
			fmt.Println("→ ", reverse(text))

		case "cap":
			fmt.Println("→ ", cap(text))

		case "title":
			fmt.Println("→ ", title(text))

		default:
			fmt.Printf("Unknown command: %q\nValid commands: upper, lower, cap, title, snake, reverse, exit\n", command)
		}
		fmt.Println()
	}
}


func main(){
	stringTransformer()
}