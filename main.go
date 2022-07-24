package main

import (
	"fmt"
	"strings"
)

func main() {
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	if command != "1232" {
		fmt.Printf("the result is %v", exit)
	}
}
