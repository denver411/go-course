package calc

import "math"

type operator struct {
	apply    func(a float64, b float64) float64
	priority uint8
}

var operators = map[string]operator{
	"+": {apply: func(a float64, b float64) float64 { return a + b }, priority: 0},
	"-": {apply: func(a float64, b float64) float64 { return a - b }, priority: 0},
	"/": {apply: func(a float64, b float64) float64 { return a / b }, priority: 1},
	"*": {apply: func(a float64, b float64) float64 { return a * b }, priority: 1},
	"^": {apply: func(a float64, b float64) float64 { return math.Pow(a, b) }, priority: 2},
}
