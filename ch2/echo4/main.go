// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "\n", "separator")
var p = flag.Int("myInt", 5, "getting an int")

func main() {
	flag.Parse()
	// *sep = "\n"
	fmt.Print(strings.Join(flag.Args(), *sep))
	// fmt.Println(*sep)
	if !*n {
		fmt.Println()
		fmt.Println(*p)
	}
}

//!-
