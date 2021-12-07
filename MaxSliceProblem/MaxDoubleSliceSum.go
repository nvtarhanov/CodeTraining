// A non-empty array A consisting of N integers is given.

// A triplet (X, Y, Z), such that 0 ≤ X < Y < Z < N, is called a double slice.

// The sum of double slice (X, Y, Z) is the total of A[X + 1] + A[X + 2] + ... + A[Y − 1] + A[Y + 1] + A[Y + 2] + ... + A[Z − 1].

// For example, array A such that:

//     A[0] = 3
//     A[1] = 2
//     A[2] = 6
//     A[3] = -1
//     A[4] = 4
//     A[5] = 5
//     A[6] = -1
//     A[7] = 2
// contains the following example double slices:

// double slice (0, 3, 6), sum is 2 + 6 + 4 + 5 = 17,
// double slice (0, 3, 7), sum is 2 + 6 + 4 + 5 − 1 = 16,
// double slice (3, 4, 5), sum is 0.
// The goal is to find the maximal sum of any double slice.

// Write a function:

// func Solution(A []int) int

// that, given a non-empty array A consisting of N integers, returns the maximal sum of any double slice.

// For example, given:

//     A[0] = 3
//     A[1] = 2
//     A[2] = 6
//     A[3] = -1
//     A[4] = 4
//     A[5] = 5
//     A[6] = -1
//     A[7] = 2
// the function should return 17, because no double slice of array A has a sum of greater than 17.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [3..100,000];
// each element of array A is an integer within the range [−10,000..10,000].

package main

// you can also use imports, for example:
//import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func MaxSumSolution(A []int) int {

	n := len(A)

	maxStart := make([]int, n)
	maxEnd := make([]int, n)
	maxSum := 0

	if n == 3 {
		return 0
	}

	for i := 1; i < n-1; i++ {
		if A[i]+maxEnd[i-1] > 0 {
			maxEnd[i] = A[i] + maxEnd[i-1]
		} else {
			maxEnd[i] = 0
		}
	}

	for i := n - 2; i > 0; i-- {
		if A[i]+maxStart[i+1] > 0 {
			maxStart[i] = A[i] + maxStart[i+1]
		} else {
			maxStart[i] = 0
		}
	}

	for i := 1; i < n-1; i++ {
		if maxEnd[i-1]+maxStart[i+1] > maxSum {
			maxSum = maxEnd[i-1] + maxStart[i+1]
		}
	}

	return maxSum

}
