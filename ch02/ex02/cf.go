package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/monmaru/ProgrammingLanguageGo/ch02/ex01/tempconv"
	"github.com/monmaru/ProgrammingLanguageGo/ch02/ex02/lengthconv"
	"github.com/monmaru/ProgrammingLanguageGo/ch02/ex02/weightconv"
)

func main() {
	var input = os.Args[1:]
	if len(input) == 0 {
		var value string
		fmt.Scan(&value)
		showAllFormat(value)
		return
	}

	for _, arg := range input {
		showAllFormat(arg)
	}
}

func showAllFormat(input string) {
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		return
	}

	showTemp(num)
	showLength(num)
	showWeight(num)
}

func showTemp(temp float64) {
	f := tempconv.Fahrenheit(temp)
	c := tempconv.Celsius(temp)
	fmt.Printf("TEMP: %s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func showLength(length float64) {
	f := lengthconv.Feet(length)
	m := lengthconv.Meter(length)
	fmt.Printf("LENGTH: %s = %s, %s = %s\n", f, lengthconv.FToM(f), m, lengthconv.MToF(m))
}

func showWeight(weight float64) {
	p := weightconv.Pound(weight)
	kg := weightconv.Kilogram(weight)
	fmt.Printf("WEIGHT: %s = %s, %s = %s\n", p, weightconv.PToKG(p), kg, weightconv.KGToP(kg))
}
