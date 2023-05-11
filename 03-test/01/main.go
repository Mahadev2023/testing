package main

import "fmt"

func main() {

	fmt.Println("sum 2 +3", mySum(2, 3))
	fmt.Println("sum 3+3", mySum(3, 3))
	fmt.Println("sum 2+2", mySum(2, 2))
}
func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
