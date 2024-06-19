package main

import (
	"app/tst/lib"
	"testing"
)

func TestBTreeSort(t *testing.T) {

	base := lib.CreateRandList(100, 100)
	bTree := lib.CreateBTree(base)

	resA := bTree.Sort(lib.ASC)

	var prev *int
	for _, v := range resA {
		if prev == nil {
			prev = &v
			continue
		}

		if v >= *prev {
			prev = &v
			continue
		}

		t.Fatalf("Fail err elem %v \n Base:%v \n Res:%v", v, base, resA)
	}
}
