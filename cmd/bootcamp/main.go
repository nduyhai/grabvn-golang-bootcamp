package main

import (
	"fmt"
)

func main() {

	numbers := make(chan int)

	go func(<-chan int) {
		for {
			num, ok := <-numbers
			if !ok {
				fmt.Println("Wtf")
				return
			}
			fmt.Println("Got:", num)
		}
	}(numbers)

	go func(chan<- int) {
		for i := 0; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}(numbers)

	_, _ = fmt.Scanln()
}
