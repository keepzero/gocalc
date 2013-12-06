package gocalc

// http://rosettacode.org/wiki/Parsing/Shunting-yard_algorithm#Go

import (
	"strings"
)

var opa = map[string]struct {
	prec   int
	rAssoc bool
}{
	"^": {4, true},
	"*": {3, false},
	"/": {3, false},
	"+": {2, false},
	"-": {2, false},
}

func parseInfix(e string) (rpn string) {
	infix := formatExpress(e)
	var stack []string // holds operators and left parenthesis
	for _, tok := range strings.Fields(infix) {
		switch tok {
		case "(":
			stack = append(stack, tok) // push "(" to stack
		case ")":
			var op string
			for {
				// pop item ("(" or operator) from stack
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // discard "("
				}
				rpn += " " + op // add operator to result
			}
		default:
			if o1, isOp := opa[tok]; isOp {
				// token is an operator
				for len(stack) > 0 {
					// consider top item on stack
					op := stack[len(stack)-1]
					if o2, isOp := opa[op]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rAssoc {
						break
					}
					// top item is an operator that needs to come off
					stack = stack[:len(stack)-1] // pop it
					rpn += " " + op              // add it to result
				}
				// push operator (the new one) to stack
				stack = append(stack, tok)
			} else { // token is an operand
				if rpn > "" {
					rpn += " "
				}
				rpn += tok // add operand to result
			}
		}
	}
	// drain stack to result
	for len(stack) > 0 {
		rpn += " " + stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return
}

func formatExpress(e string) (infix string) {
	infix = e[0:1]
	for i := 1; i < len(e); i++ {
		c := e[i]
		if in(c, "1234567890.") {
			if in(e[i-1], "1234567890.") {
				infix += string(c)
			} else {
				infix += " " + string(c)
			}
		} else if c == ' ' {
			continue
		} else {
			infix += " " + string(c)
		}
	}
	return strings.TrimSpace(infix)
}

func in(c byte, str string) bool {
	if strings.IndexByte(str, c) == -1 {
		return false
	}
	return true
}
