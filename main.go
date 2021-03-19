// From "The Go Programming Languagae by "Alan A. A. Donovan & Brian W. Kernighan

// "Dup2" prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

// Modified to print the names of all files in which each duplicated line occurs. 

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
    appearances := make(map[string]string)
    var arg string
	files := os.Args[1:]
    
	if len(files) == 0 {
		countLines(os.Stdin, counts, appearances, arg)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			//countLines(f, counts)
            countLines(f, counts, appearances, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
            fmt.Println(appearances[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, appearances map[string]string, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
        appearances[input.Text()] += arg + " "
	}
}
