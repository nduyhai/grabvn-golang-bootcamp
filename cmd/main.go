package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Golang Bootcamp Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		// convert CRLF to LF
		input = strings.Replace(input, "\n", "", -1)

		if input != "" {
			handle(input)
		}

	}

}

func handle(input string) {

	r := regexp.MustCompile("\\s+")
	result := r.Split(input, -1)

	if len(result) == 3 {
		a, error := strconv.Atoi(result[0])
		if error != nil {
			fmt.Println(a)
		}

		b, error := strconv.Atoi(result[2])
		if error != nil {
			fmt.Println(b)
		}

		switch result[1] {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		default:
			handleError(input)
		}
	} else {
		handleError(input)
	}
}

func handleError(input string) {
	fmt.Println("Sorry, please type valid input")

}
