package main

import (
	"cs-go-course5/dzhebov/ht12_hash_table/hashtable"
	"fmt"
)

/*
Реализуйте простейшую хэш таблицу с решением коллизий методом цепочек.
Алгоритм перестроения таблицы с коэффициентом заполнения строить не обязательно.

Интерфейс следующий
type HashTable interface {
	Set(key, val []byte)
	Get(key []byte) []byte
	Delete(key []byte) bool
}
*/

func main() {
	t := hashtable.MakeTable(4)
	t.Set([]byte("key1"), []byte("value1"))
	t.Set([]byte("key2"), []byte("value2"))
	t.Set([]byte("key3"), []byte("value3"))
	t.Set([]byte("key4"), []byte("value4"))
	t.Set([]byte("key5"), []byte("value5"))
	t.Set([]byte("key6"), []byte("value6"))
	t.Set([]byte("key7"), []byte("value7"))
	t.Set([]byte("key8"), []byte("value8"))
	t.Set([]byte("key9"), []byte("value9"))

	fmt.Printf("\ntable: %v", t)

	fmt.Printf("\nis key7 exist: %v", string(t.Get([]byte("key7"))[:]))
	fmt.Printf("\nis key exist: %v", t.Get([]byte("key")))

	t.Delete([]byte("key9"))
	t.Delete([]byte("key2"))
	t.Delete([]byte("key8"))
	fmt.Printf("\ntable after deleting: %v", t)
}
