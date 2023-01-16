package main

import (
	"fmt"

	"github.com/gopl.io/solution/ch5/solutions"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Max value is: %d\n", solutions.Max(mySlice))
	solutions.Min()
}
