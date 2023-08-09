package main

import "fmt"

type I[T any] interface {
	Get(id int) (T, error)
}

type S[T any] struct{}

func (s S[T]) Get(id int) (t T, err error) {
	return t, nil
}

type K struct{}

func (k K) Get(id int) (string, error) {
	return "nil", nil
}

func main() {
	var is I[string] = &S[string]{}
	fmt.Println(is.Get(1))

	var ik I[string] = &K{}
	fmt.Println(ik.Get(1))
}
