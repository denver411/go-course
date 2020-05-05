package main

import (
	"fmt"
	"os"
	"strings"

	c "./calc"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Input some expression. Only operations + - * / ^ () are supported")
		return
	}

	result, postfix := c.GetResults(args)
	fmt.Println("Преобразование: " + strings.Join(args, " ") + " --> " + strings.Join(postfix, " "))
	fmt.Println("--")
	fmt.Println("Результат: " + fmt.Sprintf("%g", result))
}
