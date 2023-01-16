package main

import (
	"fmt"

	// _ "github.com/jamesconfy/ch2/tempconv"
	_ "github.com/jamesconfy/ch2/tempconv"
	"github.com/jamesconfy/solution/ch5/5.15/solutions"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Max value is: %d\n", solutions.Max(mySlice))
	solutions.Min()
}
