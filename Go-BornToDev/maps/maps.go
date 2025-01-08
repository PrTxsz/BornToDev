package main

import "fmt"

var Product = make(map[string]float64)

func main() {
	fmt.Println("Product =", Product)

	//Add data
	Product["Macbook"] = 40000
	Product["Iphone"] = 30000
	Product["Ipad"] = 25000
	fmt.Println("Product =", Product)

	//Delete data
	delete(Product, "Ipad")
	fmt.Println("Product =", Product)

	//Update data
	Product["Iphone"] = 20000
	fmt.Println("Product =", Product)

	value1 := Product["Iphone"]
	fmt.Println("Product =", value1)

	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println("courseName =", courseName)
	fmt.Println("courseName 101 =", courseName["101"])

	//Find key by value
	FindCourse := "Java"
	for key, value := range courseName {
		if value == FindCourse {
			fmt.Println("Course =", FindCourse, "is ID", key)
		}
	}
}
