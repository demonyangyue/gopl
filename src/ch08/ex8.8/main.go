// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, s *sync.WaitGroup) {
	defer s.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	var wg sync.WaitGroup

	lines := make(chan string)

	timeout := 10 * time.Second
	timer := time.NewTimer(timeout)

	go scan(c, lines)

	// NOTE: ignoring potential errors from input.Err()
	defer func() {
		wg.Wait()
		c.Close()
	}()

	for {
		select {
		case line := <- lines:
			timer.Reset(timeout)
			wg.Add(1)
			go echo(c, line, 1*time.Second, &wg)
		case <- timer.C:
			return
		}

	}
}

func scan(c io.Reader, lines chan<- string) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		lines <- input.Text()
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
