package main

import (
	_ "golang.org/x/sync/errgroup"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ExpSum(4)
}

// 4
// 3 + 1
// 2 + 2
// 2 + 1 + 1
// 1 + 1 + 1 + 1
func ExpSum(n uint64) uint64 {
	nums := make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		nums[i] = 1
	}

	var resultMap = make(map[string]struct{})

	for i := uint64(1); i <= uint64(len(nums)); i++ {
		resultMap[createKey(elementByPartition(nums, i))] = struct{}{}
	}

	log.Println(resultMap)

	return 1
}

func createKey(nums []uint64) string {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	b := strings.Builder{}
	for _, v := range nums {
		b.WriteString(strconv.FormatUint(v, 10))
	}

	return b.String()
}

func elementByPartition(nums []uint64, partition uint64) (res []uint64) {
	var buf []uint64
	for k, i := range nums {
		buf = append(buf, i)
		if uint64(k+1)%partition == 0 {
			res = append(res, sum(buf))
			buf = []uint64{}
		}
	}
	res = append(res, buf...)
	return
}

func sum(n []uint64) (res uint64) {
	for _, v := range n {
		res += v
	}

	return
}

//
//1 + 1 + 1 + 1
//|1 + 1| + |1 + 1|
//2 + 2
//|1 + 1 + 1| + 1
//3 + 1
//1 + |1 + 1| + 1
//1 + 2 + 1
//1 + | 1 + 1 + 1 |
//1 + 3
