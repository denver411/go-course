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

func (t token) String() string {
	return t.v
}

func createToken(t tokenT, v string) token {
	return token{t: t, v: v}
}

func (t token) isOpenBracket() bool {
	return t.t == bracket && t.v == "("
}

func (ts tokens) String() string {
	res := make([]string, len(ts))

	for i, t := range ts {
		res[i] = t.String()
	}

	return strings.Join(res, " ")
}

func (ts tokens) last() token {
	return ts[len(ts)-1]
}

func (ts *tokens) pop() token {
	last := ts.last()
	*ts = (*ts)[:len(*ts)-1]
	return last
}

func (ts *tokens) push(new token) int {
	*ts = append(*ts, new)
	return len(*ts)
}

func (ts tokens) empty() bool {
	return len(ts) == 0
}
