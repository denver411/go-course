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

	postfix, err := calc.GetPostfixNotation(in)
	if err != nil {
		fmt.Println("Error during transform to postfix:", err)
		return
	}

	result, err := calc.GetPostfixResult(postfix)

	if err != nil {
		fmt.Println("Error during calculation:", err)
		return
	}

	fmt.Println("Преобразование: " + strings.Join(in, " ") + " --> " + strings.Join(postfix, " "))
	fmt.Println("--")
	fmt.Println("Результат: " + fmt.Sprintf("%g", result))
}
