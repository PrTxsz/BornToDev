package main

import "fmt"

func zeropointer(ipointer *int) {
	*ipointer = 0
}

func zerovalue(ivalue int) {
	ivalue = 0
}

func zerovaluebynewvariable(ivalue int) int {
	ivalue = 0
	return ivalue
}

func main() {
	var x int
	i := 3
	fmt.Println("i =", i)

	x = zerovaluebynewvariable(i)
	fmt.Println("i value from fn.zerovaluenewvariable =", x)

	zerovalue(i)
	fmt.Println("i value from fm.zerovalue =", i)
	// result : i value from fm.zerovalue = 3
	// because value of i in main and zerovalue is diffent function

	zeropointer(&i)
	fmt.Println("i address from fn.zeropointer =", i)
}
