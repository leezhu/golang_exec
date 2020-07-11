package main

import (
	"fmt"

	"github.com/leezhu/golang_exec/go_practice/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("olleh"))
	fmt.Printf("%T", "s")
	fmt.Println()
}

//迭代法

// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return root
// 	}
// 	queue := []*TreeNode{root}
// 	for len(queue) > 0 {
// 		size := len(queue)
// 		for i := 0; i < size; i++ {
// 			node := queue[i]
// 			node.Left, node.Right = node.Right, node.Left
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 		}
// 		queue = queue[size:]
// 	}
// 	return root
// }
