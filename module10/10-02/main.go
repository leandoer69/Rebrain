package main

import "fmt"

// пробрасываю значения с помощью -ldflags
func main() {
	var Version string
	var Name string
	fmt.Printf("Welcome! Version: %s; Name: %s", Version, Name)
}
