package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

const DEFAULT_LINES = 10

var (
	n     = flag.Int("n", DEFAULT_LINES, "")
	lines = flag.Int("lines", DEFAULT_LINES, "")
)

func main() {
	flag.Parse()
	args := flag.Args()

	num := DEFAULT_LINES
	// nとlinesが両方ある場合はnが優先される
	if *lines != DEFAULT_LINES {
		num = *lines
	}
	if *n != DEFAULT_LINES {
		num = *n
	}

	if len(args) == 0 {
		head(os.Stdin, num)
		os.Exit(0)
	}
	for _, path := range args {
		if err := handleFile(path, num); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}
}

func handleFile(path string, n int) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("head: %s: No such file or directory", path)
	}
	defer f.Close()

	head(f, n)
	return nil
}

func head(r io.Reader, n int) {
	scanner := bufio.NewScanner(r)
	for range n {
		if !scanner.Scan() {
			break
		}
		fmt.Fprintln(os.Stdout, scanner.Text())
	}
}
