package go_leetcode

import (
	"bytes"
	"math"
)

type Stage string

const (
	StageScanPrefix = "scan_prefix"
	StageScanSign   = "scan_sign"
	StageScanNum    = "scan_num"
	StageScanOther  = "scan_other"
)

func myAtoi(str string) int {

	nums := bytes.Buffer{}
	sign := 1

	stage := StageScanPrefix
	for _, c := range str {
		switch stage {
		case StageScanPrefix:
			if c == ' ' {
				continue
			}
			stage = StageScanSign
			fallthrough
		case StageScanSign:
			if c == '+' || c == '-' {
				if c == '-' {
					sign = -1
				}
				stage = StageScanNum
				continue
			}
			stage = StageScanNum
			fallthrough
		case StageScanNum:
			if c >= '0' && c <= '9' {
				nums.WriteRune(c)
				continue
			}
			stage = StageScanOther
			fallthrough
		default:
			break
		}
	}

	num := 0
	res := 0
	if nums.Len() != 0 {
		for _, v := range nums.Bytes() {
			num = num*10 + int(v-48)
			res = num * sign
			if res > math.MaxInt32 {
				return math.MaxInt32
			}
			if res < math.MinInt32 {
				return math.MinInt32
			}
		}
		return res
	}
	return res
}
