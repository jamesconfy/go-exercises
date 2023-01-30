package main

import (
	"fmt"

	"gopl.io/solution/ch5"
	"gopl.io/solution/ch7"
)

func main() {
	buf := []byte("Hello Confidence James \n James Confidence Says Hello\nWhat?")
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Max value is: %d\n", ch5.Max(mySlice...))
	fmt.Printf("Max value is: %d\n", ch5.Min(mySlice...))
	// fmt.Printf("Words is: %d\n", ch7.WordCounter(buf))
	var c ch7.WordCounter
	c.Write(buf)
	fmt.Printf("Line is: %d\n", c)

	var f ch7.LineCounter
	f.Write(buf)
	fmt.Printf("Line is: %d\n", f)
	// fmt.Printf()
}
