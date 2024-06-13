package lib

import "sort"

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

func NativeSortAdapter(base []int, order SortOrder) []int {
	sort.Slice(base, func(i, j int) bool {
		switch order {
		case ASC:
			return base[i] < base[j]
		case DESC:
			return base[i] > base[j]
		default:
			return base[i] < base[j]
		}
	})

	return base
}
