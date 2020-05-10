package calc

import (
	"fmt"
)

// Calculate function
func Calculate(expression []string) (float64, []string) {
	tokens, err := Tokenize(expression)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}
	postfix, err := GetPostfixNotation(tokens)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	result, err := GetPostfixResult(postfix)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	return result, postfix
}
