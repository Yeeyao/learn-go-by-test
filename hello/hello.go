package main

import "fmt"

//func main() {
//	fmt.Println("Hello, world")
//}

const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
