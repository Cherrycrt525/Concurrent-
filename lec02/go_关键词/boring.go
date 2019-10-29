package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		//Receive expression is just a value
		fmt.Printf("U say: %q\n", <-c)
	}
	fmt.Println("U're broing; I'm leaving.")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		//Expression to be send can be any suitable value
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
