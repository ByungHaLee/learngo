package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"me", "you"}
	for _, person := range people {
		go isNice(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isNice(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
