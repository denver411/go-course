package main

import (
	"fmt"
	"os"
	"strings"

	"./calc"
)

func main() {
	in := os.Args[1:]

	if len(in) == 0 {
		fmt.Println("Input some expression. Only operations + - * / ^ () are supported")
		return
	}

	result, postfix := calc.Calculate(in)
	fmt.Println("Преобразование: " + strings.Join(in, " ") + " --> " + strings.Join(postfix, " "))
	fmt.Println("--")
	fmt.Println("Результат: " + fmt.Sprintf("%g", result))
}
