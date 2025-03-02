package main

import (
	"fmt"
	"math"
)

// which starting number under target produces the longest collatz sequence
func Solve(target int) int {
	var longest = 0
	ans := 0
	for i := 1; i < target; i++ {
		times := collatz(i, 0)
		if times > longest {
			longest = times
			ans = i

		}
	}
	return ans
}
func collatz(num int, times int) int {
	if num == 1 {
		return times + 1
	}
	if num%2 == 0 {
		// fmt.Println("Collatz num is: ", num)
		return collatz(num/2, times+1)
	} else {
		// fmt.Println("Collatz num is: ", num)
		var temp = (3 * num)
		return collatz(temp+1, times+1)
	}

}

func main() {
	fmt.Println(Solve(int(math.Pow10(6))))
}
