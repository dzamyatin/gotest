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

func BenchmarkBTreeSort(b *testing.B) {
	b.SetBytes(185)
	b.ReportAllocs()

	base := lib.CreateRandList(100, 100)
	bTree := lib.CreateBTree(base)

	b.ResetTimer() //Do not include previous code to bench testing

	for i := 0; i < b.N; i++ {
		bTree.Sort(lib.ASC)
		//b.StopTimer() //in case if we should do some operation on each cycle iteration but we dont want to include the result in bench
		//b.StartTimer()
	}
}
