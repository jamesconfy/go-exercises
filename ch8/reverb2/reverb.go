// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

// !+
func handleConn(c net.Conn) {
	wg := sync.WaitGroup{}
	defer func() {
		wg.Wait()
		// NOTE: ignoring potential errors from input.Err()
		c.Close()
	}()
	timer := time.NewTimer(10 * time.Second)
	inputs := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			inputs <- input.Text()
		}
		if input.Err() != nil {
			log.Println("scan:", input.Err())
		}
	}()

	for {
		select {
		case input := <-inputs:
			reset := 1 * time.Second
			timer.Reset(10 * time.Second)
			wg.Add(1)
			go func() {
				defer wg.Done()
				echo(c, input, reset)
			}()
		case <-timer.C:
			by := bytes.NewBufferString("Closing gracefully\n")
			p := by.Bytes()
			c.Write([]byte(p))
			return
		}

	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
