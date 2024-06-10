package lib

type SortOrder int8

const (
	ASC SortOrder = iota
	DESC
)

func BTreeSortAdapter(base []int, order SortOrder) []int {
	bTree := CreateBTree(base)
	if bTree == nil {
		return []int{}
	}
	return bTree.Sort(order)
}

func BTreeSortAsyncAdapter(base []int, order SortOrder) []int {
	bTree := CreateBTree(base)
	if bTree == nil {
		return []int{}
	}
	return bTree.AsyncSort(order)
}
