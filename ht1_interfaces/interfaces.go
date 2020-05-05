package main

import (
	"errors"
	"fmt"
)

// Написать реализацию данного интерфейса.
// Обратить внимание, что в 2-х методах принимается любое кол-во значений.
// Не забыть при делении обработать деление на ноль.

type Calculator interface {
	Substruct(a, b interface{}) (interface{}, error)
	Sum(a ...interface{}) (interface{}, error)
	Divide(a, b interface{}) (interface{}, error)
	Multiply(a ...interface{}) (interface{}, error)
}

type Calc struct{}

func main() {
	s := Calc{}

	sub, err := s.Substruct(6, 2)
	if err != nil {
		fmt.Printf("Error during substruct: %v", err)
		return
	}
	fmt.Println(sub)

	div, err := s.Divide(6., 4.)
	if err != nil {
		fmt.Printf("Error during divide: %v", err)
		return
	}
	fmt.Println(div)

	sum, err := s.Sum(6, 5, 4)
	if err != nil {
		fmt.Printf("Error during sum: %v", err)
		return
	}
	fmt.Println(sum)

	mul, err := s.Multiply(6, 5, 10, "8")
	if err != nil {
		fmt.Printf("Error during multiply: %v", err)
		return
	}
	fmt.Println(mul)
}

func (c Calc) Substruct(a, b interface{}) (interface{}, error) {
	_, aIsInt := a.(int)
	_, bIsInt := b.(int)

	if !aIsInt || !bIsInt {
		return 0, errors.New("wrong arguments type")
	}

	return a.(int) - b.(int), nil
}

func (c Calc) Divide(a, b interface{}) (interface{}, error) {
	_, aIsInt := a.(float64)
	_, bIsInt := b.(float64)

	if !aIsInt || !bIsInt {
		return 0, errors.New("wrong arguments type")
	}
	if b == 0. {
		return 0, errors.New("divide by zero")
	}

	return a.(float64) / b.(float64), nil
}

func (c Calc) Sum(a ...interface{}) (interface{}, error) {
	if len(a) == 0 {
		return 0, errors.New("empty arguments")
	}
	s := 0

	for _, el := range a {
		switch el.(type) {
		case int:
			s += el.(int)
		default:
			return 0, errors.New("wrong arguments type")
		}
	}

	return s, nil
}

func (c Calc) Multiply(a ...interface{}) (interface{}, error) {
	if len(a) == 0 {
		return 0, errors.New("empty arguments")
	}
	m := 1

	for _, el := range a {
		switch el.(type) {
		case int:
			m *= el.(int)
		default:
			return 0, errors.New("wrong arguments type")
		}
	}

	return m, nil
}
