package main

import (
	"fmt"
)

// 2. Инициализировать слайс с элементами любого типа.
// Развернуть его (первый элемент->последний и тд).
// Сделать циклический сдвиг на N элементов
const step = 8

func main() {
	//sl := []interface{}{1, 2, 3, 4, 5, 6, 7}
	sl := []interface{}{}

	slRev := SliceReverse(sl)
	slWithStep := SliceWithStep(sl, step)
	fmt.Println("original slice: ", sl)
	fmt.Println("reversed slice: ", slRev)
	fmt.Printf("slice with step %v: %v", step, slWithStep)

}

func SliceReverse(sl []interface{}) []interface{} {
	if len(sl) == 0 {
		return make([]interface{}, 0)
	}
	var slRes []interface{}
	slRes = append(slRes, sl...)

	for i := 0; i < len(sl)/2; i++ {
		slRes[i] = sl[len(sl)-1-i]
		slRes[len(sl)-1-i] = sl[i]
	}

	return slRes
}

func SliceWithStep(sl []interface{}, step int64) []interface{} {
	if len(sl) == 0 {
		return make([]interface{}, 0)
	}

	slRes := make([]interface{}, len(sl))
	for step > int64(len(sl)) {
		step = step - int64(len(sl))
	}
	for i := 0; i < len(sl); i++ {
		idx := int64(i) + step
		if idx >= int64(len(sl)) {
			idx = idx - int64(len(sl))
		}
		slRes[idx] = sl[i]
	}

	return slRes
}
