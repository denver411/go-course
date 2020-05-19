package calc

import (
	"errors"
)

func tokenize(expression string) (tokens, error) {
	res := tokens{}
	acc := ""

	for _, ch := range expression {
		tokenType, isToken := tokensMap[ch]

		if !isToken {
			return nil, errors.New("wrong symbol")
		}

		switch tokenType {
		case separator:
			if acc != "" {
				res.push(createToken(literal, acc))
				acc = ""
			}
		case literal:
			acc = acc + string(ch)
		case bracket, operator:
			if acc != "" {
				res.push(createToken(literal, acc))
				acc = ""
			}
			res.push(createToken(tokenType, string(ch)))
		}

	}

	if acc != "" {
		res.push(token{t: literal, v: acc})
	}

	return res, nil
}
