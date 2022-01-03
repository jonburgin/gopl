package main

import (
	"fmt"
	"os"
)

func main() {
	//var s, sep string
	fmt.Println(os.Args[0])
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}
