package calc

import (
	"errors"
	"strconv"
)

// GetPostfixResult function
func GetPostfixResult(expression []string) (float64, error) {
	stack := []float64{}

	for _, char := range expression {
		operator, isOperator := operators[char]

		var a, b float64

		switch {
		case string(char) == " ":
			continue
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
