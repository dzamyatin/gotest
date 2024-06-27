package main

import (
	"fmt"
	_ "golang.org/x/sync/errgroup"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(
		ExpSum(4), //5
		ExpSum(5), //7
		ExpSum(10),
	)
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
		for j := uint64(1); j <= uint64(len(nums)); j++ {
			for t := uint64(1); t <= uint64(len(nums)); t++ {
				resultMap[createKey(elementByPartition(nums, i, j, t-1))] = struct{}{}
			}
		}
		for j := uint64(1); j <= uint64(len(nums)); j++ {
			for t := uint64(1); t <= uint64(len(nums)); t++ {
				resultMap[createKey(elementByPartition(nums, i, j, t-1))] = struct{}{}
			}
		}
	}

	log.Println(resultMap)

	return uint64(len(resultMap))
}

func createKey(nums []uint64) string {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	b := strings.Builder{}
	for _, v := range nums {
		b.WriteString(strconv.FormatUint(v, 10))
		b.WriteString(",")
	}

	return b.String()
}

func elementByPartition(
	nums []uint64,
	partition uint64,
	shift uint64,
	gap uint64,
) (res []uint64) {
	var buf []uint64
	var outbuf []uint64
	var puf uint64
	var curGap = gap
	for _, i := range nums {
		//for _, i := range nums {

		if gap == 0 || curGap == 0 {
			//if gap == 0 || gap%uint64(k+1) == 0 {
			puf++
			curGap = gap
		} else {
			outbuf = append(outbuf, i)
			curGap--
			continue
		}

		buf = append(buf, i)
		r := puf + shift + 1

		if r%partition == 0 {
			res = append(res, sum(buf))
			buf = []uint64{}
		}

		//}
	}

	res = append(res, buf...)
	res = append(res, outbuf...)
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
