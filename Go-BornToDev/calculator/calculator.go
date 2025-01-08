package main

import (
	"fmt"
	"reflect"
)

func main() {
	var input string

	//for {
	fmt.Println(" === CALCUALTOR === ")

	fmt.Print(" # Enter Number : ")
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	fmt.Println("Input datatype is", reflect.TypeOf(input))

	//}
}
