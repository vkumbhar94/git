package main

import "fmt"

var a = false

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("This is a test.")

	_ = xyz()
}

func xyz() error {
	if a {
		return nil
	}

	return fmt.Errorf("error")
}
