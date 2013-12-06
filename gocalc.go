package gocalc

// Calculate
func Calc(express string) (float64, error) {
	rpn := parseInfix(express)
	res := calcNotation(rpn)
	return res, nil
}
