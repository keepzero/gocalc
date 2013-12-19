package gocalc

import (
	"fmt"
	"strings"
)

// Calculate
func Calc(express string) (float64, error) {
	rpn := parseInfix(express)
	res := calcNotation(rpn)
	return res, nil
}

// Calc Fx
func CalcFx(fx string, m map[string]interface{}) (float64, error) {
	var express string = fx
	for str, v := range m {
		express = strings.Replace(express, str, fmt.Sprint(v), -1)
	}
	fmt.Println(express)
	return Calc(express)
}
