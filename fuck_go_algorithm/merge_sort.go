//.package 归并排序
//分冶的方法，
package main

import (
	"fmt"
)

func merge_sort(arr []int) []int {
	sort(&arr, 0, len(arr)-1)
	return arr
}

func sort(arr *[]int, left, right int) {
	if left == right { //说明元素只有一个，无需进行合并排序
		return
	}
	//取中间值
	mid := int((left + right) / 2)
	sort(arr, left, mid)         //左边
	sort(arr, mid+1, right)      //右边
	merge(arr, left, mid, right) //合并
}

func merge(arr *[]int, left, mid, right int) {
	//两个数组进行合并，使用双指针法，
	//因为两个数组都是顺序的，所以只需要比较两者大小
	//并将小的放到结果中
	var result []int = make([]int, 0)
	var p, q int = left, mid + 1
	for p <= mid && q <= right {
		if (*arr)[p] < (*arr)[q] {
			result = append(result, (*arr)[p])
			p++
		} else if (*arr)[p] > (*arr)[q] {
			result = append(result, (*arr)[q])
			q++
		} else {
			result = append(result, (*arr)[q])
			result = append(result, (*arr)[p])
			q++
			p++
		}
	}
	for p <= mid {
		result = append(result, (*arr)[p])
		p++
	}
	for q <= right {
		result = append(result, (*arr)[q])
		q++
	}
	// fmt.Println(result, " ", left, right)
	//赋值给arr
	for i := left; i <= right; i++ {
		(*arr)[i] = result[i-left]
	}
}

// func swap(a, b int) (a, b int) {
// 	a, b = b, a
// 	return a, b
// }

func main() {
	list := []int{1, 3, 4, 2, 7, 4}
	sort_list := merge_sort(list)
	fmt.Println(sort_list)
}
