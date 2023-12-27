package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	helloStr := "Hello, OTUS!"
	reversedHelloStr := reverse.String(helloStr)
	fmt.Println(reversedHelloStr)
}
