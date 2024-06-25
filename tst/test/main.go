package main

import (
	"fmt"
	_ "golang.org/x/sync/errgroup"
	"log"
	"strings"
)

//[-:T --:M ---:O -----:0 ----.:9 ---..:8 ---...:: --.:G --.-:Q --..:Z --..--:, --...:7 -.:N -.-:K -.--:Y -.--.:( -.--.-:) -.-.:C -.-.--:! -.-.-.:; -..:D -..-:X -..-.:/ -...:B -...-:= -....:6 -....-:- .:E .-:A .--:W .---:J .----:1 .----.:' .--.:P .--.-.:@ .-.:R .-.-.:+ .-.-.-:. .-..:L .-..-.:" .-...:& ..:I ..-:U ..---:2 ..--.-:_ ..--..:? ..-.:F ...:S ...-:V ...--:3 ...---...:SOS ...-..-:$ ....:H ....-:4 .....:5]

func main() {
	res := DecodeBits("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011")
	log.Println(string(res))
	log.Println(DecodeMorse(res))
	//res := DecodeBits("111100000000000000000000000000001111")
	//res := DecodeBits("110011")
	res = DecodeBits("101")
	log.Println(string(res))
	log.Println(DecodeMorse(res))
	//
	res = DecodeBits("10001")
	log.Println(string(res))
	log.Println(DecodeMorse(res))
	//
	res = DecodeBits("1110111") //.. ._.
	log.Println(string(res))
	log.Println(DecodeMorse(res))

}

var MORSE_CODE = map[string]string{
	"....": "H",
	".":    "E",
	"..":   "I",
	"-.--": "Y",
	".---": "J",
	"..-":  "U",
	"-..":  "D",
	"--":   "M",
}

func DecodeMorse(morseCode string) string {

	var buf []rune

	res := strings.Builder{}

	spaceCounter := 0
	for k, v := range morseCode {

		if v == ' ' {
			spaceCounter++
		}

		if spaceCounter == 3 {
			res.WriteString(" ")
			spaceCounter = 0
		}

		if v != ' ' {
			spaceCounter = 0
		}

		isLast := k == len(morseCode)-1

		if v == ' ' || isLast {
			if isLast {
				buf = append(buf, v)
			}

			if len(buf) > 0 {
				res.WriteString(getChar(buf))
			}

			buf = []rune{}
			continue
		}

		buf = append(buf, v)
	}

	if len(buf) > 0 {
		res.WriteString(getChar(buf))
	}

	return res.String()
}

func getChar(buf []rune) string {
	v, ok := MORSE_CODE[string(buf)]
	if ok {
		return v
	}
	panic(fmt.Sprintf(
		"there is no char in table %s '%s'",
		MORSE_CODE,
		string(buf),
	))
}

func DecodeBits(bits string) string {

	fmt.Printf(bits)

	b := []byte(bits)

	unit := getUnitSize(b)

	res := strings.Builder{}

	buf := byte(0)
	cnt := 0
	for k, v := range b {
		if k == 0 {
			buf = v
		}

		if buf != v {
			res.WriteString(convert(buf, cnt, unit))

			buf = v
			cnt = 1
			continue
		}

		cnt++
	}

	res.WriteString(convert(buf, cnt, unit))

	return res.String()
}

func convert(buf byte, cnt int, unit int) string {
	mod := cnt / unit

	if buf == '1' {
		switch mod {
		case 1:
			return "."
		case 3:
			return "-"
		}
	}

	if buf == '0' {
		switch mod {
		case 3:
			return " "
		case 7:
			return "   "
		}
	}

	return ""
}

func getUnitSize(bits []byte) int {
	min0, _, _, min1, _ := getSizeStat(bits)

	if min0 != 0 && min1 != 0 {
		if min1 < min0 {
			return min1
		}
		return min0
	}

	if min1 != 0 {
		return min1
	}

	if min0 != 0 {
		return min0
	}

	panic("fail to guess unit size")
}

func getSizeStat(bits []byte) (
	min0 int,
	mid0 int,
	max0 int,
	min1 int,
	//mid1 int,
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

				if (cnt > min0 && cnt < max0) || mid0 == 0 {
					mid0 = cnt
				}
			} else {
				if cnt < min1 || min1 == 0 {
					min1 = cnt
				}
				if cnt > max1 || max1 == 0 {
					max1 = cnt
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
