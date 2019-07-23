package main

import (
	"fmt"
)

func main() {

	numbers := make(chan int)

	go func(<-chan int) {
		for num := range numbers {
			fmt.Println("Got:", num)
		}
	}(numbers)

	go func(chan<- int) {
		for i := 0; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}(numbers)

	select {}
}
