package lib

import (
	"log"
	"sync"
)

type TreeNode struct {
	leftLeaf  *TreeNode
	rightLeaf *TreeNode
	value     int
	count     int
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		value: value,
		count: 1,
	}
}

func (t *TreeNode) GetValues() []int {
	res := make([]int, t.count)
	for k := range res {
		res[k] = t.value
	}

	return res
}

func (t *TreeNode) Sort(order SortOrder) (res []int) {

	if order == ASC {
		if t.leftLeaf != nil {
			res = append(res, t.leftLeaf.Sort(order)...)
		}

		res = append(res, t.GetValues()...)

		if t.rightLeaf != nil {
			res = append(res, t.rightLeaf.Sort(order)...)
		}
	} else {
		if t.rightLeaf != nil {
			res = append(res, t.rightLeaf.Sort(order)...)
		}

		res = append(res, t.GetValues()...)

		if t.leftLeaf != nil {
			res = append(res, t.leftLeaf.Sort(order)...)
		}
	}

	return
}

func (t *TreeNode) AsyncSort(order SortOrder) []int {

	wg := sync.WaitGroup{}

	var resA []int
	var resB []int
	if order == ASC {

		if t.leftLeaf != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()
				resA = t.leftLeaf.Sort(order)
			}()
		}

		if t.rightLeaf != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()
				resB = t.rightLeaf.Sort(order)
			}()
		}

		wg.Wait()
	} else {
		//@TODO add desc
		log.Fatal("Not implemented yet")
		//if t.rightLeaf != nil {
		//	res = append(res, t.rightLeaf.Sort(order)...)
		//}
		//
		//res = append(res, t.GetValues()...)
		//
		//if t.leftLeaf != nil {
		//	res = append(res, t.leftLeaf.Sort(order)...)
		//}
	}

	res := make([]int, 0, len(resA)+len(t.GetValues())+len(resB))

	res = append(res, resA...)
	res = append(res, t.GetValues()...)
	res = append(res, resB...)

	return res
}

func (t *TreeNode) Bypass() (res []int) {
	res = append(res, t.GetValues()...)
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
		t.count++
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
		node := NewTreeNode(v)
		if root == nil {
			root = node
		}
		root.AddLeaf(
			node,
		)
	}
	return
}
