package main

import "fmt"

func main() {

	fmt.Println("Square of 5", square(5))

}

func square(x int) int {
	return x * x
}
