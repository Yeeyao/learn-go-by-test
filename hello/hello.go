package main

import "fmt"

//func main() {
//	fmt.Println("Hello, world")
//}

const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
