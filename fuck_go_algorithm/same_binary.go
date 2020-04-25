//.
// 相同的二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSameTree(p *TreeNode, q *TreeNode) bool {
    //左右都应该相同，节点值也相同
    if p==nil && q==nil{
        return true
    }else if (p==nil ||q==nil){
        return false
    }
    right_status:=isSameTree(p.Right,q.Right)
    left_status:=isSameTree(p.Left,q.Left)
    if right_status && left_status && (p.Val == q.Val){
        return true
    }
    return false
}