package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"me", "you"}
	for _, person := range people {
		go isNice(person, c)
	}

	fmt.Println("Waiting for message")
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isNice(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is nice"
}
