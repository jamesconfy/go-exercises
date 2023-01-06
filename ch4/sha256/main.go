// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256" //!+
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strings"
)

var hashAlgorithm = flag.String("hash", "sha256", "hash algorithm to use (sha256, sha384, or sha512)")

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	arr := [40]int{1, 2}
	fmt.Printf("%x\n%x\n%t\n%T\n%x\n", c1, c2, c1 == c2, c1, arr)
	fmt.Printf("%d\n", countDifferentBits(c1, c2))

	flag.Parse()
	// Parse command-line arguments

	// Read input from standard input
	for _, v := range os.Args[1:] {
		if strings.HasPrefix(v, "--hash") {
			continue
		}

		data := []byte(v)
		switch *hashAlgorithm {
		case "sha256":
			fmt.Printf("%x\n", sha256.Sum256(data))
		case "sha384":
			fmt.Printf("%x\n", sha512.Sum384(data))
		case "sha512":
			fmt.Printf("%x\n", sha512.Sum512(data))
		default:
			log.Fatalf("invalid hash algorithm: %s", *hashAlgorithm)
		}
	}
}

func countDifferentBits(hash1, hash2 [32]byte) int {
	// Initialize a count to keep track of the number of different bits
	count := 0

	// Iterate through the two hashes, comparing each corresponding pair of bytes
	for i := 0; i < 32; i++ {
		// XOR the two bytes and count the number of set bits (1s) in the result
		count += bits.OnesCount8(hash1[i] ^ hash2[i])
	}

	return count
}

//!-
