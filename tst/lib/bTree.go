package lib

type TreeNode struct {
	leftLeaf  *TreeNode
	rightLeaf *TreeNode
	value     int
}

func (t *TreeNode) Sort(order SortOrder) (res []int) {

	if order == ASC {
		if t.leftLeaf != nil {
			res = append(res, t.leftLeaf.Sort(order)...)
		}

		res = append(res, t.value)

		if t.rightLeaf != nil {
			res = append(res, t.rightLeaf.Sort(order)...)
		}
	} else {
		if t.rightLeaf != nil {
			res = append(res, t.rightLeaf.Sort(order)...)
		}

		res = append(res, t.value)

		if t.leftLeaf != nil {
			res = append(res, t.leftLeaf.Sort(order)...)
		}
	}

	return
}

func (t *TreeNode) Bypass() (res []int) {
	res = append(res, t.value)
	if t.leftLeaf != nil {
		res = append(res, t.leftLeaf.Bypass()...)
	}

	if t.rightLeaf != nil {
		res = append(res, t.rightLeaf.Bypass()...)
	}

	return
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

func (t *TreeNode) SearchGreater(search int) (res []int) {
	if search < t.value {
		res = append(res, t.value)
	}

	if t.rightLeaf != nil {
		res = append(res, t.rightLeaf.SearchGreater(search)...)
	}
	return
}

func (t *TreeNode) SearchLower(search int) (res []int) {

	if search > t.value {
		res = append(res, t.value)
		if t.rightLeaf != nil {
			res = append(res, t.rightLeaf.SearchLower(search)...)
		}
		return
	}

	if t.leftLeaf != nil {
		res = append(res, t.leftLeaf.SearchLower(search)...)
	}
	return
}

func CreateBTree(values []int) (root *TreeNode) {
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
