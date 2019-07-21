package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
	arg := strings.Fields(input)
	if len(arg) == 3 {
		a, err := parseParamInt(arg[0])
		if err != nil {
			return
		}

		b, err := parseParamInt(arg[2])
		if err != nil {
			return
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
				e = genError("sorry, cannot  device to zero")

			} else {
				result = strconv.FormatFloat(float64(float64(a)/float64(b)), 'g', -1, 64)
			}

		default:
			e = genError("sorry, please type valid operator")
		}
	} else {
		e = genError("sorry, please type valid input")
	}
	return
}

func parseParamInt(param string) (r int, e error) {
	r, err := strconv.Atoi(param)
	if err != nil {
		e = genError("sorry, please type valid operand")
	}
	return
}

func genError(cause string) (e error) {
	e = errors.New(cause)
	return
}
