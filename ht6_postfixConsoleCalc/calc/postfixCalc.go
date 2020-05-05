package calc

import (
	"errors"
	"math"
	"strconv"
)

type operator struct {
	apply    func(a float64, b float64) float64
	priority uint8
}

var operators = map[string]operator{
	"+": {apply: func(a float64, b float64) float64 { return a + b }, priority: 0},
	"-": {apply: func(a float64, b float64) float64 { return a - b }, priority: 0},
	"/": {apply: func(a float64, b float64) float64 { return a / b }, priority: 1},
	"*": {apply: func(a float64, b float64) float64 { return a * b }, priority: 1},
	"^": {apply: func(a float64, b float64) float64 { return math.Pow(a, b) }, priority: 2},
}

func GetPostfixResult(expression []string) (float64, error) {
	stack := []float64{}

	for _, char := range expression {
		operator, isOperator := operators[char]

		var a, b float64

		switch {
		case string(char) == " ":
			break
		case isOperator:
			b, stack = stack[len(stack)-1], stack[:len(stack)-1]
			a, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, operator.apply(a, b))
		default:
			res, _ := strconv.ParseFloat(char, 64)
			stack = append(stack, res)
		}
	}

	if len(stack) > 1 {
		return 0, errors.New("Wrong expression")
	}

	return stack[len(stack)-1], nil
}
