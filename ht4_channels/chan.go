package main

import (
	"fmt"
	"sync"
)

/*
4. Создаем канал для чисел. Создаем 1000 горутин в каждой из которых в этот канал пишем какое-нибудь число.
Создаем функцию, которая принимает этот канал и читает из него все что можно и просто складывает между собой.
В main рутине нужно дождаться выполнения этой функции и получить результат сложения.
*/

func main() {
	numbers := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				n := 1
				numbers <- n
				wg.Done()
			}()
		}
		wg.Wait()
		close(numbers)
	}()

	result := sum(numbers)
	fmt.Println(result)
}

func sum(ch <-chan int) int {
	var sum int = 0

	for n := range ch {
		sum += n
	}
	return sum
}
