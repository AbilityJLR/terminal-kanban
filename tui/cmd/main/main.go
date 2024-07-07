package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")

	for {
		time.Sleep(10 * time.Second)
	}
}
