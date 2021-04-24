package main

import (
	"fmt"
)

func main() {
	const name string = "constant"
	// var phone string = "variable"
	phone := "variable" // is equal with above
	phone = "changable"
	fmt.Println(phone)
}
