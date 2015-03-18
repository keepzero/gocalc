package gocalc

import (
	"fmt"
	"strings"
)

// Calculate
func Calc(express string) (float64, error) {
	var rpn string
	if Cache {
		var ok bool
		if rpn, ok = rpnCaches[express]; !ok {
			rpn = parseInfix(express)
			rpnCaches[express] = rpn
		}
	} else {
		rpn = parseInfix(express)
	}
	return calcNotation(rpn)
}

// Calc Fx
func CalcFx(fx string, m map[string]interface{}) (float64, error) {
	var express string = fx
	for str, v := range m {
		express = strings.Replace(express, str, fmt.Sprint(v), -1)
	}
	return Calc(express)
}
