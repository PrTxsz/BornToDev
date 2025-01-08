package main

import "fmt"

func main() {
	var input int

	for {
		fmt.Print("Enter a number : ")
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if input == 123 {
			fmt.Println("Yes input =", input)
		} else if input == 0 {
			fmt.Println("----- BYE -----")
			break
		} else {
			fmt.Println("Nope input =", input)
		}
	}

}
