package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	filnames := flag.Args()
	buf := make([]byte, 1028)

	if len(filnames) == 0 {
		filnames = append(filnames, "-")
	}
	for i := 0; i < len(filnames); i++ {
		var fil *os.File
		switch filnames[i] {
		case "-":
			fil = os.Stdin
		default:
			fullpath, err := filepath.Abs(filnames[i])
			if err != nil {
				fmt.Println("Failure on filepath.Abs")
				fmt.Println(err)
				os.Exit(1)
			}
			fil, err = os.Open(fullpath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		for n, err := fil.Read(buf); err != io.EOF; n, err = fil.Read(buf) {
			os.Stdout.Write(buf[:n])
		}
	}
}
