package main

import "fmt"

/* Declare Array normal format*/
var ProduceName [4]string

func main() {
	/* Declare Array shrot format */
	Price := [4]float32{40000, 30000, 20000, 2000}

	/*Price := [4]string{}*/

	ProduceName[0] = "Macbook"
	ProduceName[1] = "Ipad"
	ProduceName[2] = "Iphone"
	ProduceName[3] = "Airpod"
	fmt.Println(ProduceName)
	fmt.Println(Price)

	for i := 0; i < len(ProduceName); i++ {
		fmt.Println(ProduceName[i], ":", Price[i])
	}

}
