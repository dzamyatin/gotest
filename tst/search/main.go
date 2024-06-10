package main

import "fmt"

func main() {
	values := []int{
		1, 6, 7, 2, 4, 4, 9, 123, 0, 6,
	}

	res := createBTree(values)

	fmt.Println(values)
	fmt.Println(res.searchLower(6))
	fmt.Println(res.searchGreater(6))
}

type TreeNode struct {
	leftLeaf  *TreeNode
	rightLeaf *TreeNode
	value     int
}

func (t *TreeNode) AddLeaf(node *TreeNode) {
	if node.value == t.value {
		return
	}

	if node.value > t.value {
		if t.rightLeaf == nil {
			t.rightLeaf = node
			return
		}
		t.rightLeaf.AddLeaf(node)
		return
	}

	if t.leftLeaf == nil {
		t.leftLeaf = node
		return
	}
	t.leftLeaf.AddLeaf(node)
}

func (t *TreeNode) searchGreater(search int) (res []int) {
	if search < t.value {
		res = append(res, t.value)
	}

	if t.rightLeaf != nil {
		res = append(res, t.rightLeaf.searchGreater(search)...)
	}
	return
}

func (t *TreeNode) searchLower(search int) (res []int) {

	if search > t.value {
		res = append(res, t.value)
		if t.rightLeaf != nil {
			res = append(res, t.rightLeaf.searchLower(search)...)
		}
		return
	}

	if t.leftLeaf != nil {
		res = append(res, t.leftLeaf.searchLower(search)...)
	}
	return
}

func createBTree(values []int) (root *TreeNode) {
	for _, v := range values {
		node := &TreeNode{value: v}
		if root == nil {
			root = node
		}
		root.AddLeaf(
			node,
		)
	}
	return
}
