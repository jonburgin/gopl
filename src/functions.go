package main

import (
	"fmt"
	"golang.org/x/net/html"
)

func max(nums ...int) (int, error) {
	if len(nums) < 1 {
		return 0, fmt.Errorf("at least one number must be provided")
	}
	max := nums[0]
	for _, val := range nums {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func min(first int, rest ...int) int {
	min := first
	for _, val := range rest {
		if val < min {
			min = val
		}
	}
	return min
}

func variadicJoin(strs ...string) string {
	sep := ""
	result := ""
	for _, s := range strs {
		result = result + sep + s
		sep = ", "
	}
	return result
}

func ElementsByTagName(doc *html.Node, names map[string]string) []*html.Node {
	var result []*html.Node
	if _, ok := names[doc.Data]; doc.Type == html.ElementNode && ok {
		result = append(result, doc)
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, ElementsByTagName(c, names)...)
	}
	return result
}

func ElementsByTagNamez(doc *html.Node, names ...string) []*html.Node {
	nameMap := make(map[string]string)
	for _, name := range names {
		nameMap[name] = ""
	}
	return ElementsByTagName(doc, nameMap)
}

func panicFunc() (result int) {
	defer func() {
		result = recover().(int)
	}()
	panic(2)
}

func main() {
	fmt.Println(max([]int{1, 3, 6, 3, 2}...))
	fmt.Println(min(1, 3, -1, 2))
	fmt.Println(variadicJoin("hi", "cow"))
	fmt.Println(variadicJoin("nobody"))
	fmt.Println(variadicJoin())
	fmt.Println(panicFunc())
}
