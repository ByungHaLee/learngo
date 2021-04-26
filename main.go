package main

import (
	"fmt"

	"github.com/ByungHaLee/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)
}
