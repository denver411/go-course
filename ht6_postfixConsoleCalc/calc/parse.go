package calc

import (
	"errors"
)

func makePostfix(infix tokens) (tokens, error) {
	stack := tokens{}
	postfix := tokens{}
	var prev *token

	for _, t := range infix {
		switch t.t {
		case literal:
			if prev != nil && prev.t == literal {
				return nil, errors.New("two literals in a row")
			}
			postfix.push(t)
		case bracket:
			if t.isOpenBracket() {
				stack.push(t)
				continue
			}
			err := handleBracket(&stack, &postfix)

			if err != nil {
				return nil, err
			}

		case operator:
			if prev != nil && prev.t == operator {
				return nil, errors.New("two operators in a row")
			}
			handleOperator(&stack, &postfix, t)
		}

		p := t
		prev = &p
	}

	for !stack.empty() {
		postfix.push(stack.pop())
	}

	return postfix, nil
}

func handleBracket(stack *tokens, res *tokens) error {
	for !stack.empty() && !stack.last().isOpenBracket() {
		last := stack.pop()
		res.push(last)
	}

	if stack.empty() {
		return errors.New("brackets not paired")
	}

	stack.pop()

	return nil
}

func handleOperator(stack *tokens, res *tokens, current token) {
	for !stack.empty() && stack.last().t == operator && isLessPriority(current.v, stack.last().v) {
		last := stack.pop()
		res.push(last)
	}
	stack.push(current)
}
