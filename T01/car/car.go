package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var params = []string{"Typecar", "Model", "Year", "Engine"}

type Car struct {
	Typecar string
	Model   string
	Year    int
	Engine  float64
}

func newCar() *Car {
	scanner := bufio.NewScanner(os.Stdin)
	c := &Car{}
	for _, v := range params {
		fmt.Print("Enter the ", v, " of your car: ")
		scanner.Scan()
		switch v {
		case "Typecar":
			c.Typecar = scanner.Text()
		case "Model":
			c.Model = scanner.Text()
		case "Year":
			c.Year, _ = strconv.Atoi(scanner.Text())
		case "Engine":
			c.Engine, _ = strconv.ParseFloat(scanner.Text(), 64)
		default:
			continue
		}
	}
	return c
}

func main() {
	c := newCar()
	fmt.Println("Your car is: ")
	jsoncar, _ := json.MarshalIndent(c, "", "  ")
	fmt.Print(string(jsoncar))
}
