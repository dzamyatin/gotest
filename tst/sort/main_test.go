package main

import (
	"app/tst/lib"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {

	sorts := []lib.SortOrder{lib.ASC, lib.DESC}

	for _, s := range sorts {
		for i := 0; i < 100; i++ {
			base := lib.CreateRandList(i, 1000)
			resA := lib.QuickSort(base, s)
			resB := lib.SlowSort(base, s)

			if resA == nil {
				resA = make([]int, 0)
			}

			if !reflect.DeepEqual(resA, resB) {
				t.Fatalf(`Error sort direction: %v
Base: %v
Res: %v
Expected: %v
`,
					s,
					base,
					resA,
					resB,
				)
			}
		}
	}
}
