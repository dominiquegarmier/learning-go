package main

import "fmt"

func quicksort(lst *[]int) {
	_quicksort(lst, 0, len(*lst))
}

func _quicksort(lst *[]int, low int, high int) {
	if (high - low) < 2 {
		return
	}
	p := partition(lst, low, high)
	_quicksort(lst, low, p)
	_quicksort(lst, p+1, high)
}

func partition(lst *[]int, low int, high int) int {
	pivot := (*lst)[high-1]
	i := low
	for j := low; j < high; j++ {
		if (*lst)[j] <= pivot {
			(*lst)[i], (*lst)[j] = (*lst)[j], (*lst)[i]
			i++
		}
	}
	(*lst)[high-1], (*lst)[i] = (*lst)[i], (*lst)[high-1]
	return i
}

func main() {
	lst := []int{10, 12, 40, 1, 2, 3, 4, 4, 1, 2, 12, 3, 1, 23, 12, 3}
	quicksort(&lst)
	fmt.Println(lst)
}
