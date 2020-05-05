package main

import (
	"fmt"
	mathRand "math/rand"
	"strconv"
	"sync"
	"time"
)

/*
Сделать структуру которая будет содержать мапу счетчиков. она должна реализовывать интерфейс
type Counters interface {
    Register(name string) // добавление нового счетчика
    Add(name string, n int) // Добавление значения для указанного счетчика
    Val(name string) // получение значения указанного счетчика
}

При этом не забыть про потокобезопасность с примитивом *sync.RWMutex
*/

type Counters interface {
	Register(name string)   // добавление нового счетчика
	Add(name string, n int) // Добавление значения для указанного счетчика
	Val(name string) int    // получение значения указанного счетчика
}

type Counts struct {
	sync.RWMutex
	data map[string]int
}

func getCounterName(idx int) string {
	return "counter_" + strconv.Itoa(idx)
}

const countersNumber = 20

func (c *Counts) Register(name string) {
	c.Lock()
	defer c.Unlock()

	if _, exist := c.data[name]; exist {
		return
	}

	c.data[name] = 0
}

func (c *Counts) Add(name string, n int) {
	c.Lock()
	defer c.Unlock()

	if _, exist := c.data[name]; !exist {
		c.data[name] = n
	}

	c.data[name] = c.data[name] + n
}

func (c *Counts) Val(name string) int {
	c.RLock()
	defer c.RUnlock()

	return c.data[name]
}

func main() {
	mathRand.Seed(time.Now().Unix())
	c := Counts{}
	wg := sync.WaitGroup{}
	c.data = make(map[string]int)

	for i := 1; i <= countersNumber; i++ {
		wg.Add(1)
		go func() {
			c.Register(getCounterName(i))
			wg.Done()
		}()
	}

	for i := 0; i < 200; i++ {
		key := mathRand.Intn(countersNumber) + 1
		wg.Add(1)
		go func() {
			c.Add(getCounterName(key), 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(c.data)

	f1 := c.Val("counter_1")
	f2 := c.Val("counter_2")
	f3 := c.Val("counter_3")
	fmt.Println(f1, f2, f3)
}
