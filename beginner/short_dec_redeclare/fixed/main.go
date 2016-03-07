package main

import "fmt"

func main() {
	x := "hello"
	newX(x)
}

func newX(message string) {
	x := "fixed!"
	fmt.Println(x)
}
