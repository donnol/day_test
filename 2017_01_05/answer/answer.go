package main

import "fmt"

func main() {
	fmt.Println(getSum(99, 999))

	arr := []int{2, 6, 3, 9, 1, 2}
	revert(arr)
	fmt.Println(arr)
	revert2(arr)
	fmt.Println(arr)

	fmt.Println(revertNum(12345))
	fmt.Println(revertNum(54321))
}

// 取水仙花数个数
func getSum(min, max int) int {
	count := 0

	for i := min; i <= max; i++ {
		nums := getBit(i)
		sum := 0
		for _, single := range nums {
			sum += single * single * single
		}
		if sum == i {
			fmt.Println(sum)
			count++
		}
	}
	return count
}

func getBit(i int) []int {
	nums := []int{}
	div := 10

	for j := 1; ; j++ {
		mod := i % div
		nums = append(nums, mod)
		if i < div {
			break
		}
		if mod == 0 {
			i /= div
		} else {
			i = (i - mod) / div
		}
	}
	// fmt.Println(nums)
	return nums
}

// 数组倒序--后面的元素大于时才交换
func revert(array []int) []int {
	arrlen := len(array)
	midIndex := arrlen / 2

	for i := 0; ; i++ {
		if i >= midIndex {
			break
		}
		if array[i] < array[arrlen-i-1] {
			array[i], array[arrlen-i-1] = array[arrlen-i-1], array[i]
		}
	}
	return array
}

func revert2(array []int) {
	arrlen := len(array)
	if arrlen-1 <= 0 {
		return
	}
	if array[0] < array[arrlen-1] {
		array[0], array[arrlen-1] = array[arrlen-1], array[0]
	}
	revert2(array[1:(arrlen - 1)])
}

// 将数字的位取出后，倒序组成数组: 12345 -> [5, 4, 3, 2, 1]
func revertNum(n int) []int {
	nums := getBit(n)
	revert(nums)
	return nums
}
