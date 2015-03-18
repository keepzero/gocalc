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

func BenchmarkCalc(b *testing.B) {
	Cache = false
	for i := 0; i < b.N; i++ {
		Calc("2*3/8")
	}
}

func BenchmarkCalcWithCache(b *testing.B) {
	Cache = true
	for i := 0; i < b.N; i++ {
		Calc("2*3/8")
	}
}

func BenchmarkCalcFx(b *testing.B) {
	Cache = false
	for i := 0; i < b.N; i++ {
		CalcFx("a*b/c", map[string]interface{}{"a": 2, "b": 3, "c": 8})
	}
}

func BenchmarkCalcFxWithCache(b *testing.B) {
	Cache = true
	for i := 0; i < b.N; i++ {
		CalcFx("a*b/c", map[string]interface{}{"a": 2, "b": 3, "c": 8})
	}
}
