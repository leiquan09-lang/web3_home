package main

import (
	"fmt"
)

// removeDuplicates 原地删除有序切片中的重复元素，返回新长度。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0                           // 慢指针：已处理区间的末尾
	for j := 1; j < len(nums); j++ { // 快指针：扫描待处理区间
		if nums[j] != nums[i] { // 发现新元素
			i++
			nums[i] = nums[j] // 压缩到前端
		}
	}
	return i + 1 // 新长度
}

func main() {

	fmt.Println(111)
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	newLen := removeDuplicates(nums)
	fmt.Println(nums[:newLen]) // 输出: [0 1 2 3 4]

}
