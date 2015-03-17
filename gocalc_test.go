package gocalc

import (
	"testing"
)

func TestCalc1(t *testing.T) {
	if res, err := Calc("(1+2)*3"); err != nil || res != 9.0 {
		t.Error("Calc test not pass.")
	} else {
		t.Log("Calc test passed.")
	}
}

func TestCalc2(t *testing.T) {
	if _, err := Calc("5/0"); err == nil {
		t.Error("Calc did not work as expected.")
	} else {
		t.Log("Calc test passed.", err)
	}
}
