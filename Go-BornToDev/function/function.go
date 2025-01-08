package main

import "fmt"

func hello() {
	fmt.Println("Hello Perfect")
}

func plus(value int, value2 int) {
	result := value + value2
	fmt.Println("Result :", result)
}

func plus2(value int, value2 int) int {
	return value + value2
}

func plus3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	hello()

	plus(1, 2)

	result := plus2(5, 5)
	fmt.Println("Result :", result)

	result2 := plus3value(5, 5, 5)
	fmt.Println("Result :", result2)
}
