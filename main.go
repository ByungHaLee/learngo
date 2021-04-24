package main

import "fmt"

func main() {
	names := []string{"aaa", "bbbb", "ccccc", "dddddd"}
	names = append(names, "eeeeeee")
	fmt.Println(names)
}
