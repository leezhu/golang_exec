//.package go_algorithm

//快速排序思想和用例
//快速排序是效率非常高的排序，复杂度在0(nlogn)。
//挖坑排序+分治思想的组合。挖坑就是我选了一个基准数值（坑），让其他数据来填，不断挖，不断填。最终
//经过一次循环后，整个数组的数据呈现出，基准数值左边都是小于它，右边则是大于它。
//分治思想就是：你既然第一次已经将数组分成了两半，那么再将左右两个数组再排序就行。依此类推，
//最终就能得到整体的排序。
//具体可以参考这一篇博文，是写的比较清晰易懂。https://www.runoob.com/w3cnote/quick-sort.html
package main

import "fmt"

func quick_sort(arr *[]int, left, right int) {
	if left >= right {
		return
	}
	l := left
	r := right
	x := (*arr)[l]
	for l < r {
		//从右边找比基准值小的数
		for l < r && (*arr)[r] >= x {
			r--
		}
		if l < r {
			(*arr)[l] = (*arr)[r] //找到后，前指针需要向后移一位
			l++
		}
		//从左边找比基准值大的数
		for l < r && (*arr)[l] < x {
			l++
		}
		if l < r {
			(*arr)[r] = (*arr)[l]
			r--
		}
		// fmt.Println("arr:", *arr, "l:", l, "r:", r)
	}
	(*arr)[l] = x               //将最后的坑由它自己填
	quick_sort(arr, left, l-1)  //左边再去找
	quick_sort(arr, l+1, right) //右边再去找
}

func main() {
	test_arr := []int{2, 5, 3, 1, 4}
	fmt.Println("begin sort:", test_arr)
	quick_sort(&test_arr, 0, 4) //传指针修改slice内容
	fmt.Println("after sort:", test_arr)
}
