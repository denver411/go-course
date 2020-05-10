package calc

import (
	"errors"
	"strconv"
	"strings"
)

// var tokens = make(map[string]string)

var tokens = map[string]string{
	"+": "+",
	"-": "-",
	"/": "/",
	"*": "*",
	"^": "^",
	"(": "(",
	")": ")",
}

func isInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}
	return true
}

// Tokenize func
func Tokenize(in []string) ([]string, error) {
	expression := strings.Join(in, "")
	res := []string{}
	temp := ""

	for _, el := range expression {
		token, isToken := tokens[string(el)]
		isDot := string(el) == "."
		isInt := isInt(string(el))
		switch {
		case isToken:
			if temp != "" {
				res = append(res, temp)
				temp = ""
			}
			res = append(res, token)

		case isInt || isDot:
			temp = temp + string(el)
		default:
			return nil, errors.New("wrong symbol in expression: " + string(el))
		}
	}

	if temp != "" {
		res = append(res, temp)
	}

	return res, nil
}
