package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)
type clock struct {
	name, host string
}

func (c clock) watch(dst io.Writer, src io.Reader) {
	s := bufio.NewScanner(src)
	for s.Scan() {
		fmt.Fprintf(dst, "%s:%s\n", c.name, s.Text())
	}

	fmt.Println(c.name, "done")
	if s.Err() != nil {
		log.Printf("can't read from %s: %s", c.name, s.Err())
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: clockwall NAME=HOST ...")
		os.Exit(1)
	}


	clocks := make([]*clock, 0)
	for _, arg := range os.Args[1:] {
		fields := strings.Split(arg, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad arg: %s\n", arg)
			os.Exit(1)
		}
		clocks = append(clocks, &clock{fields[0], fields[1]})
	}

	for _, c := range clocks {
		conn, err := net.Dial("tcp", c.host)

		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go c.watch(os.Stdout, conn)
	}

	for {
		time.Sleep(time.Minute)
	}

}

func watch(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
