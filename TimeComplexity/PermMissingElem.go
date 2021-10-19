// An array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.

// Your goal is to find that missing element.

// Write a function:

// class Solution { public int solution(int[] A); }

// that, given an array A, returns the value of the missing element.

// For example, given array A such that:

//   A[0] = 2
//   A[1] = 3
//   A[2] = 1
//   A[3] = 5
// the function should return 4, as it is the missing element.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [0..100,000];
// the elements of A are all distinct;
// each element of array A is an integer within the range [1..(N + 1)].

package main

// you can also use imports, for example:
import "sort"

func Solution(A []int) int {

	var b int

	lenA := len(A)

	if lenA == 0 {
		return 1
	}

	sort.Ints(A)

	if A[0] != 1 {
		return 1
	}

	if A[lenA-1] != lenA+1 {
		return lenA + 1
	}

	for i, value := range A {
		if (lenA != i+1) && (A[i+1]-A[i]) != 1 {
			b = value + 1
			return b
		}
	}
	return b
}
