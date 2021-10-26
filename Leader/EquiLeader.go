// A non-empty array A consisting of N integers is given.

// The leader of this array is the value that occurs in more than half of the elements of A.

// An equi leader is an index S such that 0 ≤ S < N − 1 and two sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N − 1] have leaders of the same value.

// For example, given array A such that:

//     A[0] = 4
//     A[1] = 3
//     A[2] = 4
//     A[3] = 4
//     A[4] = 4
//     A[5] = 2
// we can find two equi leaders:

// 0, because sequences: (4) and (3, 4, 4, 4, 2) have the same leader, whose value is 4.
// 2, because sequences: (4, 3, 4) and (4, 4, 2) have the same leader, whose value is 4.
// The goal is to count the number of equi leaders.

// Write a function:

// class Solution { public int solution(int[] A); }

// that, given a non-empty array A consisting of N integers, returns the number of equi leaders.

// For example, given:

//     A[0] = 4
//     A[1] = 3
//     A[2] = 4
//     A[3] = 4
//     A[4] = 4
//     A[5] = 2
// the function should return 2, as explained above.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [1..100,000];
// each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].

package main

// you can also use imports, for example:
//import "fmt"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {

	// 1. Find Lider of a slice if exists
	// else return 0
	// 2. Calculate prefix-sums
	// 3. Calculate equi leader

	equiLeaderCount := 0
	leaderCount := 0
	n := len(A)

	buffSlice := make([]int, len(A))
	prefixSums := make([]int, len(A))

	copy(buffSlice, A)

	sort.Ints(A)

	candidate := A[n/2]

	for _, value := range A {
		if value == candidate {
			leaderCount++
		}
	}

	if leaderCount < n/2 {
		//Leader not exists
		return 0
	}

	//Use prefix-sums
	leaderCount = 0
	for i, value := range buffSlice {
		if value == candidate {
			leaderCount++

		}
		prefixSums[i] = leaderCount
	}

	//Find equi-leaders
	for i := 0; i < n-1; i++ {
		if prefixSums[i] != 0 {
			//Maybe we can calculate equi-liders withou this condition
			if i == 0 {
				if (prefixSums[n-1] - prefixSums[0]) > (n-1)/2 {
					equiLeaderCount++
				}
			} else {
				if (prefixSums[i] > (i+1)/2) &&
					((prefixSums[n-1] - prefixSums[i]) > (n-(i+1))/2) {
					equiLeaderCount++
				}
			}

		}
	}
	return equiLeaderCount
}
