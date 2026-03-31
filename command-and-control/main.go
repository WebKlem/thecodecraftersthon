// // CodeCrafters — Hackathon 002
// // Squad: Gophers
// // Members: David Abraham, Obeko Eunice, Ugwu Chioma, Michael Bulus, Emmanuel Elaigwu, Akatu Worthy, Samuel Jireh, Alexander Ogwu, Eyikwodani clementina

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func baseConverter() {
	var hex, dec, bin string

	var base int
base:
	fmt.Println("1: DECIMAL TO HEXADECIMAL AND BINARY.")
	fmt.Println("2: HEXADECIMAL TO DECIMAL")
	fmt.Println("3: BINARY TO DECIMAL")
	fmt.Print("4: GO TO MENU \n \n")

	fmt.Print("SELECT BASE TO BE CONVERTED: ")
	fmt.Scanln(&base)
	fmt.Println()

	for {
		switch base {
		case 1:
		decimalnumber:
			fmt.Print("ENTER DECIMAL NUMBER: ")
			fmt.Scanln(&dec)

			decN, err := strconv.ParseInt(dec, 10, 64)
			if err != nil {
				fmt.Println()
				fmt.Println("Invalid Decimal Number")
				fmt.Print("Enter a valid Decimal Number \n \n")
				goto decimalnumber
			}
			hexN := strconv.FormatInt(decN, 16)
			binN := strconv.FormatInt(decN, 2)

			fmt.Printf("The Binary conversion of %q is %v\n", dec, binN)
			fmt.Printf("The HexaDecimal conversion of %q is %v\n \n", dec, strings.ToUpper(hexN))

			fmt.Println("Do you have another Decimal Number to convert?")
			fmt.Println(" 1. Yes\n 2. No, Go to Base\n")
			var choice int
			fmt.Scan(&choice)

			if choice == 1 {
				goto decimalnumber
			} else if choice == 2 {
				goto base
			} else if choice < 1 || choice > 2 {
				fmt.Println("Chhose from the option")
			}

		case 2:

		hexadecimalnumber:
			fmt.Print("ENTER HEX NUMBER: ")
			fmt.Scanln(&hex)
			dec, err := strconv.ParseInt(hex, 16, 64)
			if err != nil {
				fmt.Println()
				fmt.Println("Invalid Hexadecimal Number")
				fmt.Print("Enter a valid Hexadecimal Number \n \n")
				goto hexadecimalnumber
			}
			fmt.Println(dec)
			continue

		case 3:

		binarynumber:
			fmt.Print("ENTER BINARY NUMBER: ")
			fmt.Scanln(&bin)

			dec, err := strconv.ParseInt(bin, 2, 64)
			if err != nil {
				fmt.Println("Invalid Binary Number ")
				fmt.Print("Enter a valid Binary Number \n \n")
				goto binarynumber
			}
			fmt.Println(dec)
			continue

		case 4:
			fmt.Print("Returning to Menu...\n \n")
			return

		default:
			fmt.Println("Invalid Option")
			fmt.Print("Choose A Valid Option \n \n")
			goto base
		}

	}
}

func calculator() {
	var history [5]string
	var index int = 0
	var lastResult float64
	for {
		var a float64
		var b float64

		fmt.Println("...GOPHER'S CALC...")
	firstNumber:
		fmt.Println("Input first number")
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Print("Enter digit only!\n")
			goto firstNumber
		}

	secondNumber:
		fmt.Println("Input second number")
		fmt.Scanln(&b)
		if err != nil {
			fmt.Print("Enter digit only!\n")
			goto secondNumber
		}

		var Operator string

		fmt.Println()
		fmt.Println("add")
		fmt.Println("sub")
		fmt.Println("mul")
		fmt.Println("div")
		fmt.Println("remainder")
		fmt.Println("power")
		fmt.Println("last")
		fmt.Println("history")
		fmt.Println("Exit")
		fmt.Println()
		fmt.Print("Type: ")
		fmt.Scanln(&Operator)

		switch Operator {
		case "add":
			result := a + b
			fmt.Println(result)

			entry := fmt.Sprintf("%g + %g = %g", a, b, result)
			history[index%5] = entry
			index++
			lastResult = result
			continue

		case "sub":
			result := a - b
			fmt.Println(result)

			entry := fmt.Sprintf("%g - %g = %g", a, b, result)
			history[index%5] = entry
			index++
			lastResult = result
			continue

		case "mul":
			result := a * b
			fmt.Println(result)

			entry := fmt.Sprintf("%g * %g = %g", a, b, result)
			history[index%5] = entry
			index++
			lastResult = result
			continue

		case "div":
			if b == 0 {
				fmt.Println("Not divisible by zero")
			}
			result := a / b
			fmt.Println(result)

			entry := fmt.Sprintf("%g / %g = %g", a, b, result)
			history[index%5] = entry
			index++
			lastResult = result
			continue

		case "remainder":
			result := int64(a) % int64(b)
			fmt.Println(result)

			entry := fmt.Sprintf("%d %% %d = %d", int64(a), int64(b), result)
			history[index%5] = entry
			index++
			lastResult = float64(result)
			continue

		case "power":
			result := math.Pow(a, b)
			fmt.Println(result)

			entry := fmt.Sprintf("%g ^ %g = %g", a, b, result)
			history[index%5] = entry
			index++
			lastResult = result
			continue

		case "last":
			fmt.Println("Last Result: ", lastResult)
			continue

		case "history":
			fmt.Println("Last 5 Calculationns: ")
			for i := 0; i < 5; i++ {
				if history[i] != "" {
					fmt.Println(history[i])
				}
			}
			continue

		case "Exit":
			fmt.Println("Exiting...")
			break

		default:
			fmt.Print("Out of operation range\n", "Choose from the above\n")
			continue
		}
		break
	}
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

func history() {
	fmt.Println()
	fmt.Println("STILL WORKING ON IT")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

start:
	fmt.Print("════════════════════════════════════════════════\n  THE GOPHER'S SENTINEL — COMMAND & CONTROL CONSOLE\n     All systems nominal. Type 'help' to begin.\n════════════════════════════════════════════════\nC&C>\n \n")
	fmt.Println("     calc   <command>   → the calculator")
	fmt.Println("     base   <command>   → the base converter")
	fmt.Println("     str    <command>   → the string transformer")
	fmt.Println("     help               → shows all commands")
	fmt.Println("     history            → shows last 10 inputs")
	fmt.Print("     exit               → shuts down the console\n \n")

	fmt.Print("Type Here: ")

	scanner.Scan()
	input := scanner.Text()

	for {

		switch input {
		case "calc":
			calculator()
			goto start

		case "base":
			baseConverter()
			goto start

		case "str":
			stringTransformer()
			goto start

		case "help":

			fmt.Println()
			fmt.Print("The calculator: \n  Used for arithemetic problems.\n  Type 'calc' to execute the calculator.  \n \n The base converter: \n  Used to convert from one base to another. \n  Type 'base' to execute the base converter. \n \n The string transformer: \n  Used for transforming strings. \n  Type 'str' to execute string transformer \n \n")
			goto start

		case "history":
			history()
			fmt.Println()
			goto start

		case "exit":
			fmt.Print("Goodbye! \nExiting... \n")
			break

		default:
			fmt.Print("Type the correct command \n \n")
			goto start
		}
		break
	}

}
