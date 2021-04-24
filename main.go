package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"burger", "ramen"}
	p := person{name: "aa", age: 11, favFood: favFood}
	fmt.Println(p)

}
