//prints the numbers from 1 to 100
//For multiples of 3 ==> print "Fizz" instead of the num
//For 5 ==> print "Buzz"
//For both 3 & 5 ==> "fizzBuzz"

package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
