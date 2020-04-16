//镜像二叉树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
*/
var res [][]int //存层级遍历二维数组
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res = [][]int{}
	res = append(res, make([]int, 0))
	res[0] = append(res[0], root.Val)
	depth := 1
	doLevelTraverse(root, depth)
	//迭代检查是否对称
	if len(res) == 0 {
		return true
	}
	// return true
	fmt.Println(res, len(res))
	for _, v := range res {
		// fmt.Println(v)
		if checkMirror(v) == false {
			return false
		}
	}
	return true
}

//获取层级遍历的二维数组
func doLevelTraverse(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	// fmt.Println("depth=",depth," len(res)=",len(res))
	if len(res) < (depth + 1) { //小于层级，那么新建一级
		res = append(res, make([]int, 0))
	}
	if root.Left != nil {
		res[depth] = append(res[depth], root.Left.Val)
		doLevelTraverse(root.Left, depth+1)
		//如果右边为空，那么需要补上
		// if root.Right==nil{
		//     res[depth]=append(res[depth],-1,-1)
		// }
	} else {
		//存空值
		res[depth] = append(res[depth], -1)
	}
	if root.Right != nil {
		res[depth] = append(res[depth], root.Right.Val)
		doLevelTraverse(root.Right, depth+1)
		//如果左边为空，那么需要补上
		// if root.Left==nil{
		//     res[depth]=append(res[depth],-1,-1)
		// }
	} else {
		//存空值
		res[depth] = append(res[depth], -1)
	}
}

//检查是否对称
func checkMirror(arr []int) bool {
	i := 0
	j := len(arr) - 1
	for i < j {
		// fmt.Println("i=",i," j=",j,"arr=",arr)
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}