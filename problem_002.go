/*
https://projecteuler.net/problem=2



Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

*/

package main

import "fmt"

func main() {
	curr := 1
	next := 1
	var tmpNext int

	sum := 0

	for next <= 4_000_000 {
		fmt.Printf("%d ", next)
		if next%2 == 0 {
			sum += next
		}

		tmpNext = next
		next += curr
		curr = tmpNext
	}

	fmt.Printf("\nAns is: %d\n", sum)
}
