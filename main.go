package main

import (
	"fmt"
)

func main() {

	A := []int{3, 1, 2, 4, 3}
	fmt.Println(Solution(A))

}

func Solution(A []int) int {

	//Dont do that in real project
	minSum := 99999999999
	leftSum := A[0]

	var sliceSum int
	for _, value := range A {
		sliceSum += value
	}

	for i := range A {

		if i == 1 {
			rightSum := sliceSum - A[0]
			sum := getAbsValue(leftSum - rightSum)
			if sum < minSum {
				minSum = int(sum)
			}
		} else if i > 1 && i != 0 {
			leftSum += A[i-1]
			rightSum := sliceSum - leftSum
			sum := getAbsValue(leftSum - rightSum)
			if sum < minSum {
				minSum = int(sum)
			}
		}
	}

	return minSum
}

func getAbsValue(v int) int {
	if v < 0 {
		v *= -1
	}

	return v
}
