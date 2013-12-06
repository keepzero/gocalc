package gocalc

// http://rosettacode.org/wiki/Parsing/RPN_calculator_algorithm#Go

import (
	"math"
	"strconv"
	"strings"
)

func calcNotation(rpn string) float64 {

	var stack []float64
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
			f, _ := strconv.ParseFloat(tok, 64)
			stack = append(stack, f)

		}
	}
	return stack[0]
}
