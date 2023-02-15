// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 231.

// Pipeline3 demonstrates a finite 3-stage pipeline
// with range, close, and unidirectional channel types.
package main

import "fmt"

// !+
func counter(send chan<- int) {
	for x := 0; x < 100; x++ {
		send <- x
	}
	close(send)
}

func squarer(send chan<- int, receive <-chan int) {
	for v := range receive {
		send <- v * v
	}
	close(send)
}

func printer(receive <-chan int) {
	for v := range receive {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

//!-
