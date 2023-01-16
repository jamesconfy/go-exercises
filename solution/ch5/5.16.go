package ch5

// func main() {
// 	var list = []string{"James", "Confidence", "Is", "A", "Boy"}
// 	fmt.Printf("Joined list is: %s\n", Join(list...))
// }

func Join(vals ...string) string {
	var result string

	for _, val := range vals {
		result += val
		result += "\t"
	}

	return result
}
