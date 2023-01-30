// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package ch7

import (
	"bufio"
	"bytes"
)

//!+wordcounter

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	count := 0
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

//!-wordcounter

//!+linecounter

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	count := 0
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}
