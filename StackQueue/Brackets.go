package main

// A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:

// S is empty;
// S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
// S has the form "VW" where V and W are properly nested strings.
// For example, the string "{[()()]}" is properly nested but "([)()]" is not.

// Write a function:

// func Solution(S string) int

// that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

// For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [0..200,000];
// string S consists only of the following characters: "(", "{", "[", "]", "}" and/or ")".

func Solution(S string) int {

	if S == "" {
		return 1
	}

	n := 0
	topValueOfStack := ""
	stringValue := ""

	buffStack := []string{}

	for _, value := range S {

		stringValue = string(value)

		if stringValue == "(" || stringValue == "[" || stringValue == "{" {
			// Push into stack
			buffStack = append(buffStack, stringValue)
			n++
		}

		if stringValue == ")" || stringValue == "]" || stringValue == "}" {

			if n == 0 {
				return 0
			}

			topValueOfStack = buffStack[n-1]
			if ((topValueOfStack == "(") && (stringValue == ")")) ||
				((topValueOfStack == "[") && (stringValue == "]")) ||
				((topValueOfStack == "{") && (stringValue == "}")) {
				// Pop operation
				buffStack = buffStack[:n-1]
				n--
			} else {
				return 0
			}
		}

	}
	if n == 0 {
		return 1
	} else {
		return 0
	}
}
