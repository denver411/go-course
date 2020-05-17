package calc

import "strings"

type tokenT uint8

type token struct {
	t tokenT
	v string
}

type tokens []token

const (
	operator  tokenT = 0
	literal   tokenT = 1
	bracket   tokenT = 2
	separator tokenT = 3
)

var tokensMap = map[rune]tokenT{
	'+': operator,
	'-': operator,
	'/': operator,
	'*': operator,
	'^': operator,
	' ': separator,
	'(': bracket,
	')': bracket,
	'.': literal,
	'0': literal,
	'1': literal,
	'2': literal,
	'3': literal,
	'4': literal,
	'5': literal,
	'6': literal,
	'7': literal,
	'8': literal,
	'9': literal,
}

func (t token) toString() string {
	return t.v
}

func createToken(t tokenT, v rune) token {
	return token{t: t, v: string(v)}
}

func isOpenBracket(t token) bool {
	return t.t == bracket && t.toString() == "("
}

// TokensToString func
func TokensToString(ts tokens) string {
	res := make([]string, 0, len(ts))

	for _, t := range ts {
		res = append(res, t.toString())
	}

	return strings.Join(res, " ")
}
