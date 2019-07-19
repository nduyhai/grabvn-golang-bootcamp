package main

import (
	"bufio"
	"errors"
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
			result, e := handle(input)
			if e == nil {
				fmt.Println(result)
			} else {
				fmt.Println(e.Error())
			}
		}

	}

}

func handle(input string) (result string, e error) {

	r := regexp.MustCompile("\\s+")
	arg := r.Split(input, -1)

	if len(arg) == 3 {
		a, error := strconv.Atoi(arg[0])
		if error != nil {
			e = errors.New("sorry, please type valid operand 1")
		}

		b, error := strconv.Atoi(arg[2])
		if error != nil {
			e = errors.New("sorry, please type valid operand 2")
		}

		switch arg[1] {
		case "+":
			result = strconv.FormatInt(int64(a+b), 10)
		case "-":
			result = strconv.FormatInt(int64(a-b), 10)

		case "*":
			result = strconv.FormatInt(int64(a*b), 10)

		case "/":
			if b == 0 {
				e = errors.New("sorry, cannot  device to zero")

			} else {
				result = strconv.FormatFloat(float64(float64(a)/float64(b)), 'g', -1, 64)
			}

		default:
			e = errors.New("sorry, please type valid operator")
		}
	} else {
		e = errors.New("sorry, please type valid input")
	}
	return
}
