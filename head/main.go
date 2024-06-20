package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s n\n", os.Args[0])
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: n should be number\n")
		os.Exit(1)
	}
	if len(os.Args) == 2 {
		head(os.Stdin, n)
		os.Exit(0)
	}
	for _, path := range os.Args[2:] {
		if err := handleFile(path, n); err != nil {
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
