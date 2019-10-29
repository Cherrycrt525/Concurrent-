package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	mid := len(s) / 2

	//Make 2 new channels
	leftSum := make(chan int)
	rightSum := make(chan int)

	//Start 2 new goroutines
	go sum(s[:mid], leftSum)
	go sum(s[mid:], rightSum)

	//receive th results from the 2 goroutines
	x := <-leftSum
	y := <-rightSum

	fmt.Println(x, y, x+y)
}
