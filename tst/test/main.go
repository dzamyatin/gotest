package main

import (
	_ "golang.org/x/sync/errgroup"
	"log"
	"strings"
)

func main() {
	res := DecodeBits("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011")
	log.Println(string(res))
}

func DecodeBits(bits string) string {
	b := []byte(bits)

	unit := getUnitSize(b)

	res := strings.Builder{}
	//var res []uint16

	buf := byte(0)
	cnt := 0
	for k, v := range b {
		if k == 0 {
			buf = v
		}

		isLast := k == (len(b) - 1)

		if buf != v || isLast {
			if isLast {
				cnt++
			}

			mod := cnt / unit
			if buf == '1' {
				switch mod {
				case 1:
					res.WriteString(".")
				case 3:
					res.WriteString("−")
				default:
					//res.WriteString(fmt.Sprintf("X(%d)", mod))
				}
			}

			if buf == '0' {
				switch mod {
				case 3:
					res.WriteString(" ")
				case 7:
					res.WriteString("   ")
				default:
					//res.WriteString(fmt.Sprintf("Y(%d)", mod))
				}
			}

			buf = v
			cnt = 1
			continue
		}

		cnt++
	}

	return res.String()
}

func getUnitSize(bits []byte) int {
	min0, max0, min1, mid1, max1 := getSizeStat(bits)

	if min0 != 0 && max0 != 0 {
		return min0
	}

	if min1 != 0 && max1 != 0 && mid1 != 0 {
		return min1
	}

	panic("fail to guess unit size")
}

func getSizeStat(bits []byte) (
	min0 int,
	max0 int,
	min1 int,
	mid1 int,
	max1 int,
) {
	buf := byte(0)
	cnt := 0
	for k, v := range bits {
		if k == 0 {
			buf = v
		}

		isLast := k == len(bits)-1

		if buf != v || isLast {
			if isLast {
				cnt++
			}

			if buf == '0' {
				if cnt < min0 || min0 == 0 {
					min0 = cnt
				}

				if cnt > max0 || max0 == 0 {
					max0 = cnt
				}
			} else {
				if cnt < min1 || min1 == 0 {
					min1 = cnt
				}
				if cnt > max1 || max1 == 0 {
					max1 = cnt
				}
				if (cnt > min1 && cnt < max1) || mid1 == 0 {
					mid1 = cnt
				}
			}

			cnt = 1
			buf = v
			continue
		}

		cnt++
	}

	return
}
