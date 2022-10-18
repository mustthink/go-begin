package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var queries = []string{"Enter a: ", "Enter b: ", "Enter operation: "}
var scanner = bufio.NewScanner(os.Stdin)

func handlErr() {
	fmt.Println("Error input")
}

func inVal(q int) string {
	fmt.Print(queries[q])
	scanner.Scan()
	val := scanner.Text()

	if strings.ToLower(val) == "exit" {
		os.Exit(0)
	}
	return val
}

func calculation(a, b int, o string) float64 {
	switch strings.ToLower(o) {
	case "+", "add", "plus":
		return float64(a + b)
	case "-", "subtract", "minus":
		return float64(a - b)
	case "*", "multiply":
		return float64(a * b)
	case "/", "divide", ":":
		return float64(a) / float64(b)
	default:
		fmt.Println("Error operation")
		main()
	}
	return 0
}

//go run main.go
//enter "Exit" to exit from program

func main() {
	a, err := strconv.Atoi(inVal(0))
	if err != nil {
		handlErr()
		main()
	}
	b, err := strconv.Atoi(inVal(1))
	if err != nil {
		handlErr()
		main()
	}
	oper := inVal(2)
	fmt.Println(a, " ", oper, " ", b, " = ", calculation(a, b, oper), "\n")
	main()
}
