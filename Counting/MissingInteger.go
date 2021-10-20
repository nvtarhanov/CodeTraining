package main

// Write a function:

// func Solution(A []int) int

// that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

// For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

// Given A = [1, 2, 3], the function should return 4.

// Given A = [−1, −3], the function should return 1.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [1..100,000];
// each element of array A is an integer within the range [−1,000,000..1,000,000].

import "sort"

//It is just MVP.I Spended 20 mins on this task, needs to refactor
// Maybe should try to use Map, but it is required 2 loops
func Solution(A []int) int {

	max := 0
	lenA := len(A)
	minimumFromOthers := 0

	sort.Ints(A)

	Contains1 := false

	for i, value := range A {

		if A[i] == 1 {
			Contains1 = true
		}

		if i < lenA-1 {
			if (A[i] >= 0) && (A[i+1] >= 0) {
				if (A[i] != A[i+1]) && ((A[i+1] - A[i]) != 1) {
					minimumFromOthers = A[i] + 1
					break
				}
			}
		}

		if value > max {
			max = value
		}
	}

	if !Contains1 || max == 0 {
		return 1
	}

	if minimumFromOthers != 0 {
		return minimumFromOthers
	}

	return max + 1

}
