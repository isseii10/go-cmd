package main

import (
	"fmt"
	"io"
	"os"
)

// $ cat a b c > out
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: file name not given\n", os.Args[0])
	}
	for _, v := range os.Args[1:] {
		cat(v)
	}
}

func cat(s string) {
	f, err := os.Open(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file %s: %v", s, err)
		os.Exit(1)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file %s: %v", s, err)
		os.Exit(1)
	}
	if _, err := io.WriteString(os.Stdout, string(b)); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write: %v", err)
		os.Exit(1)
	}
}
