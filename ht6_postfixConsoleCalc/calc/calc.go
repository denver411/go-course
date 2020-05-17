package calc

import (
	"errors"
	"strconv"
)

func Calculate(tokens tokens) (float64, error) {
	stack := []float64{}

	for _, t := range tokens {
		switch t.t {
		case literal:
			f, err := strconv.ParseFloat(t.v, 64)
			if err != nil {
				return 0, err
			}

			stack = append(stack, f)
		case operator:
			if len(stack) < 2 {
				return 0, errors.New("Wrong expression")
			}
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			result, err := applyOperator(t.v, b, a)

			if err != nil {
				return 0, err
			}

			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		}
	}

	if len(stack) > 1 {
		return 0, errors.New("Wrong expression")
	}

	return stack[len(stack)-1], nil
}
