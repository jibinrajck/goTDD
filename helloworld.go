package main

import "fmt"

func main() {
	fmt.Print(hello("world"))
}

func hello(name string) string {

	return ("Hello " + name)
}
