package main

import "fmt"

// 只出现一次”的数字有且仅有一个
func single(arr []int) int {
	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}
	for num, cnt := range countMap {
		if cnt == 1 {
			return num
		}
	}
	return 0 // 兜底，虽然题目保证不会走到
}

// 回文数
func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}
	return x == rev || x == rev/10
}

func main() {
	arr := []int{4, 4, 1, 4, 6, 6}
	fmt.Println(single(arr))         // 1
	fmt.Println(isPalindrome(12321)) // true
}
