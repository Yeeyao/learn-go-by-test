package main

import "fmt"

//func main() {
//	fmt.Println("Hello, world")
//}

const helloPrefix = "Hello, "

func Hello(name string) string {
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
