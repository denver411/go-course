package main

import (
	"cs-go-course5/dzhebov/ht6_postfixConsoleCalc/calc"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := os.Args[1:]

	if len(in) == 0 {
		fmt.Println("Input some expression. Only operations + - * / ^ () are supported")
		return
	}

	expression := strings.Join(in, " ")

	tokens, err := calc.Tokenize(expression)
	if err != nil {
		fmt.Printf("error while getting postfix: %v\n", err)
		return
	}
	postfix, err := calc.Parse(tokens)
	if err != nil {
		fmt.Printf("error while getting postfix: %v\n", err)
		return
	}

	postfixString := calc.TokensToString(postfix)

	result, err := calc.Calculate(postfix)
	if err != nil {
		fmt.Printf("error while calculate: %v\n", err)
		return
	}

	fmt.Printf("Преобразование: %s --> %s", expression, postfixString)
	fmt.Printf("\n--\n")
	fmt.Printf("Результат: %g\n", result)
}
