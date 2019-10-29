package main

import (
	"fmt"
	"time"
)

func say(something string) {
	fmt.Println(something)
}

func main() {
	//Start a new goroutine running the func say()
	go say("Hello" + "World")
	time.Sleep(1 * time.Second)
}
