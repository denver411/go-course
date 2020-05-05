package calc

import "errors"

// function GetPostfixNotation
func GetPostfixNotation(infix []string) ([]string, error) {
	stack := []string{}
	postfix := []string{}

	for _, char := range infix {
		operator, isOperator := operators[char]

		switch {
		case string(char) == " ":
			break
		case char == "(":
			stack = append(stack, char)
		case char == ")":
			stack, postfix = handleBracket(stack, postfix)

			if stack == nil {
				return nil, errors.New("Wrong infix expression. Check operators. Only + - * / ^ () are supported")
			}
		case isOperator:
			stack, postfix = handleOperator(stack, postfix, char, operator)
		default:
			postfix = append(postfix, char)
		}
	}

	if len(stack) > 0 {
		for idx := range stack {
			postfix = append(postfix, stack[len(stack)-idx-1])
		}
	}

	return postfix, nil
}

func handleBracket(stack []string, res []string) ([]string, []string) {
	stackCopy := append(stack[:0:0], stack...)
	resCopy := append(res[:0:0], res...)
	for j := len(stackCopy) - 1; j >= 0; j-- {
		last := stackCopy[len(stackCopy)-1]

		if last != "(" && j != 0 {
			resCopy = append(resCopy, last)
			stackCopy = stackCopy[:len(stackCopy)-1]
		} else if last == "(" && j == 0 {
			return nil, nil
		} else if last == "(" {
			stackCopy = stackCopy[:len(stackCopy)-1]
			break
		}
	}

	return stackCopy, resCopy
}

func handleOperator(stack []string, res []string, operator string, currentOperator operator) ([]string, []string) {
	stackCopy := append(stack[:0:0], stack...)
	resCopy := append(res[:0:0], res...)

	for j := len(stackCopy) - 1; j >= 0; j-- {
		lastStackOperator, isLastOperator := operators[stackCopy[len(stackCopy)-1]]

		if isLastOperator && lastStackOperator.priority >= currentOperator.priority {
			resCopy = append(resCopy, stackCopy[len(stackCopy)-1])
			stackCopy = stackCopy[:len(stackCopy)-1]
		} else {
			break
		}
	}
	stackCopy = append(stackCopy, operator)

	return stackCopy, resCopy
}
