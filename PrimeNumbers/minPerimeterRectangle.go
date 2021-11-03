// An integer N is given, representing the area of some rectangle.

// The area of a rectangle whose sides are of length A and B is A * B, and the perimeter is 2 * (A + B).

// The goal is to find the minimal perimeter of any rectangle whose area equals N. The sides of this rectangle should be only integers.

// For example, given integer N = 30, rectangles of area 30 are:

// (1, 30), with a perimeter of 62,
// (2, 15), with a perimeter of 34,
// (3, 10), with a perimeter of 26,
// (5, 6), with a perimeter of 22.
// Write a function:

// func Solution(N int) int

// that, given an integer N, returns the minimal perimeter of any rectangle whose area is exactly equal to N.

// For example, given an integer N = 30, the function should return 22, as explained above.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [1..1,000,000,000].

package solution

import "math"

func Solution(N int) int {

	perimeter := 9999999999
	temp := 0

	if N == 1 {
		return 4
	}

	for i := 1; i*i < N; i++ {
		if N%i == 0 {
			temp = 2 * (i + N/i)
			if temp < perimeter {
				perimeter = temp
			}
		}
	}

	sqrt := int(math.Sqrt(float64(N)))
	if sqrt*sqrt == N {
		temp = sqrt * 4
		if temp < perimeter {
			perimeter = temp
		}
	}

	return perimeter

}
