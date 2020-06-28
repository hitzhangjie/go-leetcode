package go_leetcode

import "bytes"

type Stage string

const (
	StageScanPrefix = "scan_prefix"
	StageScanSign   = "scan_sign"
	StageScanNum    = "scan_num"
	StageScanOther  = "scan_other"
)

func myAtoi(str string) int {

	nums := bytes.Buffer{}
	sign := '+'

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
				sign = c
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
	if nums.Len() != 0 {
		for _, v := range nums.Bytes() {
			num = num*10 + int(v-48)
		}
		if sign == '-' {
			num = num * -1
		}
	}
	return num
}
