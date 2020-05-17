package calc

import (
	"errors"
)

// Tokenize func
func Tokenize(expression string) (tokens, error) {
	res := tokens{}
	acc := ""

	for _, char := range expression {
		tokenType, isToken := tokensMap[char]

		if !isToken {
			return nil, errors.New("wrong symbol")
		}

		switch tokenType {
		case separator:
			if acc != "" {
				res.push(token{t: literal, v: acc})
				acc = ""
			}
		case literal:
			acc = acc + string(char)
		case bracket, operator:
			if acc != "" {
				res.push(token{t: literal, v: acc})
				acc = ""
			}
			res.push(createToken(tokenType, char))
		}

	}

	if acc != "" {
		res.push(token{t: literal, v: acc})
	}

	return res, nil
}
