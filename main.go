package main

import (
	"eda/sorting"
	"fmt"
)

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}

	sorting.InsertionSort(arr)

	fmt.Println(arr)

}
