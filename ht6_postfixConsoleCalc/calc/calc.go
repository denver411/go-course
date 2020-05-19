package calc

import (
	"errors"
	"strconv"
)

// Calculate func
func Calculate(exp string) (float64, error) {
	tokens, err := tokenize(exp)
	if err != nil {
		return 0, err
	}
	stack := stackFloat{}

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
			a, b := stack.pop(), stack.pop()
			result, err := applyOperator(t.v, b, a)

			if err != nil {
				return 0, err
			}

			stack.push(result)
		}
	}

	if len(stack) > 1 {
		return 0, errors.New("Wrong expression")
	}

	return stack[len(stack)-1], nil
}

// Parse func
func Parse(exp string) (string, error) {
	tokens, err := tokenize(exp)
	if err != nil {
		return "", err
	}

	result, err := makePostfix(tokens)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
