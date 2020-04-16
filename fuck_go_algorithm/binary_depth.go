//.package fuck_go_algorithm

//最大深度,自顶向下。能往下走，就+1.更新全局的深度
package fuck_go_algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
var (
	maxmiu_depth=0
)
func max_depth(root *TreeNode,depth int) {
	if root==nil
		return 
	if root.Left==nil && root.Right{
		maxmiu_depth=math.Max(float(maxmiu_depth),float(depth))
	}
	//左节点，
	max_depth(root.Left, depth+1)
	//右节点
	max_depth(root.Right, depth+1)
}

//自底向上
func max_depth_end_to_top(root *TreeNode)(depth int){
	if root==nil
		return 0
	left_depth:=max_depth_end_to_top(root.Left)
	right_depth:=max_depth_end_to_top(root.Right)
	return math.Max(left_depth, right_depth)+1
}
