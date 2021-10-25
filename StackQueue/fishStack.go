package main

// you can also use imports, for example:

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {

	if len(A) == 1 {
		return 1
	}

	fishStack := []int{}

	n := 0

	for i := 0; i < len(B); i++ {

		if len(fishStack) == 0 {
			//PUSH for first fish
			fishStack = append(fishStack, i)
			n++
		} else {

			for j := n - 1; j >= 0; j-- {

				if (B[fishStack[j]] == B[i]) || (B[fishStack[j]] == 0 && B[i] == 1) {
					// Same or opposite directions PUSH
					fishStack = append(fishStack, i)
					n++
					break
				} else if B[fishStack[j]] == 1 && B[i] == 0 {
					if A[fishStack[j]] > A[i] {
						// Kill A[i] because of less weight
						break
					} else {
						//Kill one from stack because of weight POP
						fishStack = fishStack[:j]
						n--
						if j == 0 {
							// Check if our fish is John Rambo PUSH
							fishStack = append(fishStack, i)
							break
						}
					}
				}
			}
		}
	}
	return len(fishStack)
}
