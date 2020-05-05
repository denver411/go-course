package calc

import (
	"fmt"
)

// GetResults function
func GetResults(expression []string) (float64, []string) {
	postfix, err := GetPostfixNotation(expression)
	if err != nil {
		fmt.Println("Wrong infix expression. Gaps required. Only + - * / ^ () are supported")
		return 0, nil
	}

	result, err := GetPostfixResult(postfix)
	if err != nil {
		fmt.Println("Error while getting result")
		return 0, nil
	}

	return result, postfix
}
