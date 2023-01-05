// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
		fmt.Printf("  %s\n", commaByte(os.Args[i]))
		fmt.Printf("  %s\n", reverse(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	//integer, _ := strconv.Atoi(s)
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaByte(s string) string {
	var buf bytes.Buffer
	n := len(s)
	m := n % 3
	// fmt.Println(m)
	if m != 0 {
		fmt.Fprintf(&buf, "%v", s[:m])
		buf.WriteByte(',')
	}

	for i := m; i < n; i += 3 {
		fmt.Fprintf(&buf, "%v", s[i:i+3])
		if i+3 < n {
			buf.WriteByte(',')
		}
	}

	return buf.String()
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

//!-
