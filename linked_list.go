package main

import "errors"

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
