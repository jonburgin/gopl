package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//var counts map[string]int
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		counts["Stdin"] = -1
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, err := fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				if err != nil {
					log.Default().Println("Error formatting")
				}
				continue
			}
			counts[arg] = -1
			countLines(f, counts)
			err2 := f.Close()
			if err2 != nil {
				//ignore
			}
		}
	}
	for line, count := range counts {
		var fileName string
		if count == -1 {
			fileName = line
		}
		if count > 1 {
			fmt.Printf("From %s: %d\t%s\n", fileName, count, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) map[string]int {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	return counts
}
