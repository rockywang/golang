package mylib

import (
	"fmt"
)

/**
* 斐波那契函数
 */
func Fibonacci(value int) {
	x := make([]int, value)
	x[0], x[1] = 1, 1
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2]
	}
	for _, term := range x {
		fmt.Printf("%v\n", term)
	}

}
