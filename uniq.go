package main

import (
	"bytes"
	"os"
)

func uniq(c chan []byte, out chan []byte) {
	var prev []byte
	for curr := range c {
		if bytes.Compare(curr, prev) != 0 {
			prev = curr
			out <- curr
		}
	}
	close(out)
}

func find_elements(c chan []byte) {
	var prev []byte
	buf := make([]byte, 1028)
	for n, err := os.Stdin.Read(buf); err == nil; n, err = os.Stdin.Read(buf) {
		prev = bytes.Join([][]byte{prev, buf[:n]}, []byte{})
		for bytes.Contains(prev, []byte{'\n'}) {
			i := bytes.IndexRune(prev, '\n')
			c <- prev[:i+1]
			prev = prev[i+1:]
		}
	}
	close(c)
}

func main() {
	c := make(chan []byte, 128)
	out := make(chan []byte, 10)
	go find_elements(c)
	go uniq(c, out)

	for elem := range out {
		os.Stdout.Write(elem)
	}

}
