package main

import "fmt"

func main() {
	var num1, num2 int
	for {
		fmt.Println("enter a number: ")
		fmt.Scan(&num1)
		fmt.Println("enter another number: ")
		fmt.Scan(&num2)

		fmt.Println("select an operator")
		fmt.Println("1 - add")
		fmt.Println("2 - sub")
		fmt.Println("3 - mul")
		fmt.Println("4 - div")
		fmt.Println("5 - Exit")
		fmt.Println("6 - help")

		var input int
		fmt.Scan(&input)

		if input == 6 {
			fmt.Println("1 - add")
			fmt.Println("2 - sub")
			fmt.Println("3 - mul")
			fmt.Println("4 - div")
			fmt.Println("5 - Exit")
			fmt.Println("6 - help")

		}

		if input == 5 {
			fmt.Println("Thank you...")
			break

		}

		switch input {
		case 1:
			fmt.Println(num1 + num2)
		case 2:
			fmt.Println(num1 - num2)
		case 3:
			fmt.Println(num1 * num2)
		case 4:
			if num2 != 0 {
				fmt.Println(num1 / num2)

			} else {
				fmt.Println("error!")
			}

		}
	}
}
