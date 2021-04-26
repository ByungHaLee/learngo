package main

import (
	"fmt"
	"time"
)

func main() {
	niceCount("me")
	niceCount("you")
}

func niceCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is nice", i)
		time.Sleep(time.Second)
	}
}
