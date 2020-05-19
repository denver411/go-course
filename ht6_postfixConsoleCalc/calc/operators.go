package calc

import (
	"errors"
	"math"
)

var priorities = map[string]uint8{
	"+": 0,
	"-": 0,
	"/": 1,
	"*": 1,
	"^": 2,
}

func applyOperator(operator string, operand1 float64, operand2 float64) (float64, error) {
	var res float64 = 0

	switch operator {
	case "+":
		res = operand1 + operand2
	case "-":
		res = operand1 - operand2
	case "/":
		if operand2 == 0 {
			return 0, errors.New("Division by zero")
		}
		res = operand1 / operand2
	case "*":
		res = operand1 * operand2
	case "^":
		res = math.Pow(operand1, operand2)
	}

	return res, nil
}

func isLessPriority(operator1 string, operator2 string) bool {
	return priorities[operator2] >= priorities[operator1]
}
