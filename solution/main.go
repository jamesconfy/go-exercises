package main

import (
	"fmt"

	"gopl.io/solution/ch5"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Max value is: %d\n", ch5.Max(mySlice...))
	fmt.Printf("Max value is: %d\n", ch5.Min(mySlice...))
	// fmt.Printf()
}
