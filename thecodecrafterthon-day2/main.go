package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var input string
		var choice int

		fmt.Println("select base sytem: ")
		fmt.Println("1 = Binary")
		fmt.Println("2 = Decimal")
		fmt.Println("3 = HexDecimal")
		fmt.Println("4 = quit")
		fmt.Scan(&choice)

		fmt.Println("input number")
		fmt.Scan(&input)

		switch choice {
		case 1:
			for _, ch := range input {
				if ch != '0' && ch != '1' {
					fmt.Println("invalid bin num")
					continue
				}
			}

		case 2:
			for _, ch := range input {
				if ch < '0' || ch > '9' {
					fmt.Println("invalid decimal")
					continue
				}

			}

		case 3:
			for _, ch := range input {
				if !(ch >= '0' && ch <= '9' || ch >= 'a' && ch <= 'f' || ch >= 'A' && ch <= 'F') {
					fmt.Println("invalid Hex")
					continue

				}
			}

		case 4:
			if choice == 4 {
				fmt.Println("thank you muah")
				return
			}
		default:
			fmt.Println("Invalid base")
			continue
		}

		var Decimal int64
		var err error

		switch choice {
		case 1:
			Decimal, err = strconv.ParseInt(input, 2, 64)

		case 2:
			Decimal, err = strconv.ParseInt(input, 10, 64)

		case 3:
			Decimal, err = strconv.ParseInt(input, 16, 64)
		}

		if err != nil {
			fmt.Println("error", err)
			continue
		}

		fmt.Println("conversions:")
		fmt.Println("binary: ", strconv.FormatInt(Decimal, 2))
		fmt.Println("Decimal: ", strconv.FormatInt(Decimal, 10))
		fmt.Println("Hex: ", strconv.FormatInt(Decimal, 16))

	}
}
