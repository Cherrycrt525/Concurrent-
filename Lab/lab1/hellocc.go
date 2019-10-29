//Conccurrent -- Goroutine
//uses a for loop to start 5 goroutines
//from goroutine i where i is the num of the goroutine

package main

import (
	"fmt"
	"time"
)

//To start a new goroutine: go someFunc()
func hello(i int) {
	fmt.Println("Hello from goroutine", i)
}
func main() {
	for i := 0; i < 5; i++ {
		go hello(i)
	}
	time.Sleep(1 * time.Second)
}



