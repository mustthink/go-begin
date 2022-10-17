package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

type Car struct {
	Typecar string
	Model   string
	Year    int
	Engine  float64
}

func main() {
	var t = flag.String("typecar", "Roadster", "type of your car")
	var m = flag.String("model", "Chevy", "model of your car")
	var y = flag.Int("year", 2019, "year of your car")
	var e = flag.Float64("engine", 1998.7, "V engine")
	flag.Parse()
	c := &Car{*t, *m, *y, *e}
	fmt.Println("Your car is: ")
	jsoncar, _ := json.MarshalIndent(c, "", "  ")
	fmt.Print(string(jsoncar))
}
