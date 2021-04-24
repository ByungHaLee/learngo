package main

import (
	"fmt"

	"github.com/ByungHaLee/learngo/something"
)

func main() {
	fmt.Println("Hello World!")
	// something.sayHello() // private
	something.SayBye() // public
}
