// Question-1.go
// Write a function that takes two arrays as input, each array contains a list of A-Z; Your program should return True if the
// 2nd array is a subset of 1st array, or False if not.
// For example:
// isSubset([A,B,C,D,E], [A,E,D]) = true
// isSubset([A,B,C,D,E], [A,D,Z]) = false
// isSubset([A,D,E], [A,A,D,E]) = true
// Please explain the computational complexity of your answer in Big-O notation, i.e. O(log n) or O(n ^2)?
// m := make(map[string]string)
// request

// Answers : O(n + 1) or O(2n)

package main

import "fmt"

func main() {
	A := []string{"A", "B", "C", "D", "E"}
	// B := []string{"A", "D", "Z"} // false
	// B := []string{"A", "E", "D"} // true
	B := []string{"A", "A", "D", "E"} // true
	result := isSubset(A, B)
	fmt.Println(result)
}

func isSubset(A []string, B []string) bool {

	stack := map[string]string{}

	for i := 0; i < len(A); i++ {
		stack[A[i]] = A[i]
	}

	for k := range B {
		if _, ok := stack[B[k]]; !ok {
			return false
		}
	}

	return true
}
