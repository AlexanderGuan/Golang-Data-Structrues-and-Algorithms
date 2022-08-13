package main

import (
	"strconv"
	"strings"
)

type Node struct {
	data   int
	parent *Node
	left   *Node
	right  *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) InsertItem(i int) {
	if tree.root == nil {
		tree.root = &Node{data: i}
		return
	}
	currentNode := tree.root
	for {
		if i > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &Node{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = &Node{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.left
		}
	}
}

func (tree *BinaryTree) SearchItem(i int) (*Node, bool) {
	if tree.root == nil {
		return nil, false
	}
	currencyNode := tree.root
	for currencyNode != nil {
		if i == currencyNode.data {
			return currencyNode, true
		} else if i > currencyNode.data {
			currencyNode = currencyNode.right
		} else if i < currencyNode.data {
			currencyNode = currencyNode.left
		}
	}
	return nil, false
}

/**
* 前序遍历
 */
func (tree *BinaryTree) PreorderTraversal(root *Node, callback func(int)) {
	if root == nil {
		return
	}
	callback(root.data)
	tree.PreorderTraversal(root.left, callback)
	tree.PreorderTraversal(root.right, callback)
}

/**
* 中序遍历
 */
func (tree *BinaryTree) InorderTraversal(root *Node, callback func(int)) {
	if root == nil {
		return
	}
	tree.InorderTraversal(root.left, callback)
	callback(root.data)
	tree.InorderTraversal(root.right, callback)
}

/**
* 后序遍历
 */
func (tree *BinaryTree) PostorderTraversal(root *Node, callback func(int)) {
	if root == nil {
		return
	}
	tree.PostorderTraversal(root.left, callback)
	tree.PostorderTraversal(root.right, callback)
	callback(root.data)
}

/**
* 层序遍历BFS
 */
func LevelorderTraversal(root *Node) [][]int {
	//定义一个存储结果的二维切片
	var res [][]int
	//定义一个含有根节点的队列
	arr := []*Node{root}
	//根节点为空直接返回空
	if root == nil {
		return res
	}
	//当队列为空就返回结果，不为空就进入循环
	for len(arr) > 0 {
		//重置队列长度
		size := len(arr)
		//定义一个切片来存储出队列的值
		curRes := []int{}
		//遍历队列
		for i := 0; i < size; i++ {
			//临时变量存储队列中的元素
			node := arr[i]
			//将队列中元素的值加入到这个切片里。
			curRes = append(curRes, node.data)
			//寻找队列里元素的左右孩子，如果不为空就入队
			if node.left != nil {
				arr = append(arr, node.left)
			}
			if node.right != nil {
				arr = append(arr, node.right)
			}
		}
		//遍历队列结束后，把上一层对应的元素出队。
		arr = arr[size:]
		//把所有结果加入到最后的结果上
		res = append(res, curRes)
	}
	return res
}

/**
* 最大深度
 */
func (tree *BinaryTree) Max(num1 int, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
func (tree *BinaryTree) MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return tree.Max(tree.MaxDepth(root.left), tree.MaxDepth(root.right)) + 1 // 树的深度是root的高度，而root的高度是 左右孩子中较大者+1
}

/**
* 最近公共祖先
 */
func (tree *BinaryTree) lowestCommonAncestor(root, p, q *Node) *Node {
	if root == nil {
		return nil
	}
	if root.data == p.data || root.data == q.data {
		return root
	}

	var findLeft = tree.lowestCommonAncestor(root.left, p, q)
	var findRight = tree.lowestCommonAncestor(root.right, p, q)

	if findLeft != nil && findRight != nil {
		return root
	} else if findLeft != nil {
		return findLeft
	} else {
		return findRight
	}
}

/**
* 构造二叉树（根据前序遍历和中序遍历）
 */
func (tree *BinaryTree) buildTree1(preorder []int, inorder []int) *Node {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	left := tree.findRootIndex1(preorder[0], inorder)
	root := &Node{
		data:  preorder[0],
		left:  tree.buildTree1(preorder[1:left+1], inorder[:left]),
		right: tree.buildTree1(preorder[left+1:], inorder[left+1:])}
	return root
}
func (tree *BinaryTree) findRootIndex1(target int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}

/**
* 构造二叉树（根据后序遍历和中序遍历）
 */
func (tree *BinaryTree) buildTree2(inorder []int, postorder []int) *Node {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	//先找到根节点（后续遍历的最后一个就是根节点）
	nodeValue := postorder[len(postorder)-1]
	//从中序遍历中找到一分为二的点，左边为左子树，右边为右子树
	left := tree.findRootIndex2(inorder, nodeValue)
	//构造root
	root := &Node{
		data:  nodeValue,
		left:  tree.buildTree2(inorder[:left], postorder[:left]), //将后续遍历一分为二，左边为左子树，右边为右子树
		right: tree.buildTree2(inorder[left+1:], postorder[left:len(postorder)-1])}
	return root
}
func (tree *BinaryTree) findRootIndex2(inorder []int, target int) (index int) {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}

/**
* 二叉树的序列化与反序列化
 */
// Serializes a tree to a single string.
func (this *BinaryTree) serialize(root *Node) string {
	sb := &strings.Builder{}
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.data))
		sb.WriteByte(',')
		dfs(node.left)
		dfs(node.right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *BinaryTree) deserialize(data string) *Node {
	sp := strings.Split(data, ",")
	var build func() *Node
	build = func() *Node {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &Node{
			data:  val,
			left:  build(),
			right: build()}
	}
	return build()
}
