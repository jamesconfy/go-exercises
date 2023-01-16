package solutions

import (
	"math"
)

// func main() {
// 	mySlice := []int{1, 2, 3, 4, 5, 6, 8, 9, 11, 200, -50}
// 	fmt.Printf("Max value is: %d\n", max(mySlice...))
// 	fmt.Printf("Min value is: %d\n", min(mySlice...))
// 	fmt.Println("Working")
// }

func Max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}

	maxVal := int(math.Inf(-1))
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func Min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}

	minVal := 1000000000000000000
	// fmt.Println(minVal)
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
		// fmt.Println(minVal)
	}
	return minVal
}
