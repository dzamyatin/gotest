package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {

	sorts := []SortOrder{ASC, DESC}

	for _, s := range sorts {
		for i := 0; i < 100; i++ {
			base := createRandList(i, 1000)
			resA := quickSort(base, s)
			resB := slowSort(base, s)

			if resA == nil {
				resA = make([]int, 0)
			}

			if !reflect.DeepEqual(resA, resB) {
				t.Fatalf(`Error sort direction: %v
Base: %v
Res: %v
`,
					s,
					base,
					resA)
			}
		}
	}
}
