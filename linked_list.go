package main

import (
	"errors"
	"fmt"
)

type llnode struct {
	value int
	next  *llnode
}

type llist struct {
	first *llnode
	size  uint
}

// out of struct definitions for struct methods
func (r *llist) get(index uint) (int, error) {
	if index >= r.size {
		return 0, errors.New("index out of range")
	}
	current := r.first
	for i := uint(0); i < index; i++ {
		current = current.next
	}
	return current.value, nil
}

func (r *llist) prepend(value int) {
	r.first = &llnode{value, r.first}
	r.size++
}

func main() {
	r := llist{}
	for i := 0; i < 10; i++ {
		r.prepend(i * 2)
	}

	val, err := r.get(5)

	fmt.Println(val, err)
	fmt.Println(r.get(9))
	fmt.Println(r.get(10))
}
