package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	answer := gcd(x, y)
	fmt.Printf("Lowest common divisor: %d\n", answer)

	myMap := []string{"gold", "silver", "bronze"}
	myMap = append(myMap, "violet")
	myMap = append(myMap, "indigo")
	myMap = append(myMap, "purple")
	myMap = append(myMap, "blue")
	// sorted := []int{1, 2, 3, 4, 5, 11, 7}
	// sort.Ints(sorted)
	// fmt.Println(sorted)
	sort.SliceStable(myMap, func(i, j int) bool {
		return myMap[i] > myMap[j]
	})
	fmt.Println(myMap)
}
