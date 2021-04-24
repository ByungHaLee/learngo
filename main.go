package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	len, upper := lenAndUpper("Hello")
	fmt.Println(len, upper)
}
