package main

import "fmt"

func main() {
	person := map[string]string{"name": "name", "age": "11"}
	for key, value := range person {
		fmt.Println(key, value)
	}
	fmt.Println(person)
}
