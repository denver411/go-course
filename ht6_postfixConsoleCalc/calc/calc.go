package calc

import (
	"fmt"
)

// Calculate function
func Calculate(expression []string) (float64, []string) {
	postfix, err := GetPostfixNotation(expression)
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
