package main

import "fmt"

func simple(af func(a, b int) int) {
	fmt.Println(af(60, 7))
}

func main() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}
