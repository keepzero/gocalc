package gocalc

// http://rosettacode.org/wiki/Parsing/RPN_calculator_algorithm#Go

import (
	"math"
	"strconv"
	"strings"
)

func calcNotation(rpn string) (float64, error) {

	var stack []float64
	var err error
	for _, tok := range strings.Fields(rpn) {
		switch tok {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "^":
			stack[len(stack)-2] =
				math.Pow(stack[len(stack)-2], stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		default:
			f, err := strconv.ParseFloat(tok, 64)
			if err != nil {
				return float64(0), err
			}
			stack = append(stack, f)

		}
	}
	return stack[0], err
}
