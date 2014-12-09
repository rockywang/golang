package main

import (
	"fmt"
)

func testParam(x int, y int) {
	var big, low int
	if x > y {
		low = y
		big = x
	} else {
		low = x
		big = y
	}
	fmt.Printf("%d,%d\n", low, big)
}

type stack struct {
	i    int
	data [10]int
}

/**
* 堆栈的操作
**/
func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}

/**
* 堆栈pop
 */
func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

/**
* 实现golang中通用的map函数
*
**/
type e interface{}

func multMore(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string) + f.(string) + f.(string)
	}
	return f
}
func Map(n []e, f func(e) e) []e {
	m := make([]e, len(n))
	for k, v := range n {
		m[k] = f(v)
	}
	return m
}
