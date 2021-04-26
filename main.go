package main

import (
	"fmt"
	"time"
)

func main() {
	go niceCount("me")
	go niceCount("you")
	time.Sleep(time.Second * 5)
}

func niceCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is nice", i)
		time.Sleep(time.Second)
	}
}
