package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world! The current date is:",
		time.Now().Format("02.01.2006 15:04"))
}
