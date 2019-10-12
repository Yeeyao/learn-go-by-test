package main

import "fmt"

const testchar = "a"

func For(times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += testchar
	}
	return repeated
}

func main() {
	fmt.Println(For(5))
}
