package calc

import (
	"errors"
)

// GetPostfixNotation function
func GetPostfixNotation(infix []string) ([]string, error) {
	stack := Stack{}
	postfix := Stack{}

	for _, char := range infix {
		operator, isOperator := operators[char]

		switch {
		case string(char) == " ":
			continue
		case char == "(":
			stack.Push(char)
		case char == ")":
			var err error
			err = handleBracket(&stack, &postfix)

			if err != nil {
				return nil, err
			}
		case isOperator:
			handleOperator(&stack, &postfix, char, operator)
		default:
			postfix.Push(char)
		}
	}

	if len(stack) > 0 {
		for idx := range stack {
			postfix.Push(stack[len(stack)-idx-1])
		}
	}

	return postfix, nil
}

func handleBracket(stack *Stack, res *Stack) error {
	for j := stack.Len() - 1; j >= 0; j-- {
		last := stack.Last()

		if last != "(" && j != 0 {
			res.Push(last)
			stack.Pop()
		} else if last != "(" && j == 0 {
			return errors.New("Wrong infix expression. Check operators. Only + - * / ^ () are supported")
		} else if last == "(" {
			stack.Pop()
			break
		}
	}

	return nil
}

func handleOperator(stack *Stack, res *Stack, operator string, currentOperator operator) {
	for j := stack.Len() - 1; j >= 0; j-- {
		lastStackOperator, isLastOperator := operators[stack.Last()]

		if isLastOperator && lastStackOperator.priority >= currentOperator.priority {
			res.Push(stack.Last())
			stack.Pop()
		} else {
			break
		}
	}
	stack.Push(operator)
}
