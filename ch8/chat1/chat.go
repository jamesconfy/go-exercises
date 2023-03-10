// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// !+broadcaster
type client struct {
	Out  chan<- string // an outgoing message channel
	Name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.Out <- msg
			}

		case cli := <-entering:
			cli.Out <- "Current Presents:"
			for c := range clients {
				cli.Out <- c.Name
			}
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		}
	}
}

//!-broadcaster

// !+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{ch, who}
	ch <- time.Now().Format(time.RFC3339) + ": You are " + who
	messages <- time.Now().Format(time.RFC3339) + ": " + who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- time.Now().Format(time.RFC3339) + ": " + who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- time.Now().Format(time.RFC3339) + ": " + who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

// !+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
