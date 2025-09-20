package main

import "fmt"

func pushOn(x []int) []int {
	num := 0
	for _, d := range x {
		num = num*10 + d
	}
	num++
	s := fmt.Sprintf("%d", num)
	res := make([]int, len(s))
	for i, v := range s {
		fmt.Println(s)
		res[i] = int(v - '0')
	}
	return res
}

func main() {
	// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
	fmt.Println(pushOn([]int{9}))
}
