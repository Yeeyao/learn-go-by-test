package main

import (
	"fmt"
	"io"
	"net/http"
)

// 这里为了通用，同时因为 os.Stdout 以及 bytes.Buffer 都实现了 io.Writer 接口，所以直接使用它
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterRHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterRHandler))
}
