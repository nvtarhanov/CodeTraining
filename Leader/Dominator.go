package main

import "sort"

func Solution(A []int) int {

	//1. Sort slice
	//2. Domintaor would be in a midle of slice
	n := len(A)

	if n == 0 {
		return -1
	}

	if n == 1 {
		return 0
	}

	sort.Ints(A)
	count := 0
	ind := -1

	candidate := A[n/2]

	for i, value := range A {

		if value == candidate {
			count++
			ind = i
		}
	}

	if count > n/2 {
		return ind
	}

	return -1
}
